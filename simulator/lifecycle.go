package simulator

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

const mma2BinaryEnv = "MMA2_BINARY"

type mma2Process interface {
	Stop() error
	Done() <-chan error
}

type mma2ProcessFactory func(configPath string) (mma2Process, error)

// MMA2Activator is the structural runtime boundary used after ownership-safe
// effective configuration composition.
type MMA2Activator interface {
	Activate(configPath string, cfg EffectiveMMA2Config) error
	Stop() error
}

// MMA2Lifecycle owns exactly one MMA2 child process and replaces it only after
// the new effective configuration has been composed successfully.
type MMA2Lifecycle struct {
	mu              sync.Mutex
	factory         mma2ProcessFactory
	readyTimeout    time.Duration
	current         mma2Process
	activeConfig    []byte
	activeListeners []string
}

func NewMMA2Lifecycle() *MMA2Lifecycle {
	return &MMA2Lifecycle{factory: startMMA2Process, readyTimeout: 3 * time.Second}
}

func (m *MMA2Lifecycle) Activate(configPath string, cfg EffectiveMMA2Config) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("read effective MMA2 config: %w", err)
	}
	listeners := effectiveListeners(cfg)
	oldProcess := m.current
	oldConfig := append([]byte(nil), m.activeConfig...)
	oldListeners := append([]string(nil), m.activeListeners...)
	if oldProcess != nil {
		if err := oldProcess.Stop(); err != nil {
			return fmt.Errorf("stop active MMA2: %w", err)
		}
		m.current = nil
	}

	if len(listeners) == 0 {
		m.activeConfig = configBytes
		m.activeListeners = nil
		return nil
	}
	process, err := m.factory(configPath)
	if err == nil {
		err = waitMMA2Ready(process, listeners, m.readyTimeout)
	}
	if err == nil {
		m.current = process
		m.activeConfig = configBytes
		m.activeListeners = listeners
		return nil
	}
	if process != nil {
		_ = process.Stop()
	}

	activationErr := err
	if oldProcess != nil && len(oldConfig) > 0 {
		rollbackPath := filepath.Join(filepath.Dir(configPath), ".rollback.yaml")
		if writeErr := os.WriteFile(rollbackPath, oldConfig, 0o600); writeErr != nil {
			return fmt.Errorf("activate MMA2: %v; write rollback config: %w", activationErr, writeErr)
		}
		rollback, startErr := m.factory(rollbackPath)
		if startErr == nil {
			startErr = waitMMA2Ready(rollback, oldListeners, m.readyTimeout)
		}
		_ = os.Remove(rollbackPath)
		if startErr != nil {
			if rollback != nil {
				_ = rollback.Stop()
			}
			return fmt.Errorf("activate MMA2: %v; restore prior MMA2: %w", activationErr, startErr)
		}
		m.current = rollback
		m.activeConfig = oldConfig
		m.activeListeners = oldListeners
	}
	return fmt.Errorf("activate MMA2: %w", activationErr)
}

func (m *MMA2Lifecycle) Stop() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.current == nil {
		return nil
	}
	err := m.current.Stop()
	m.current = nil
	return err
}

func effectiveListeners(cfg EffectiveMMA2Config) []string {
	listeners := make([]string, 0, len(cfg.Listeners))
	for _, listener := range cfg.Listeners {
		if listener.Listen != "" {
			listeners = append(listeners, listener.Listen)
		}
	}
	return listeners
}

func waitMMA2Ready(process mma2Process, listeners []string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for {
		allReady := true
		for _, address := range listeners {
			host, port, err := net.SplitHostPort(address)
			if err != nil {
				return fmt.Errorf("invalid MMA2 listener %q: %w", address, err)
			}
			if host == "" || host == "0.0.0.0" || host == "::" {
				host = "127.0.0.1"
			}
			conn, dialErr := net.DialTimeout("tcp", net.JoinHostPort(host, port), 75*time.Millisecond)
			if dialErr != nil {
				allReady = false
				break
			}
			_ = conn.Close()
		}
		if allReady {
			// A different process may already own the configured socket. Require
			// the child to remain alive briefly after every listener is reachable.
			select {
			case err := <-process.Done():
				if err == nil {
					err = errors.New("process exited during readiness confirmation")
				}
				return err
			case <-time.After(100 * time.Millisecond):
				return nil
			}
		}
		select {
		case err := <-process.Done():
			if err == nil {
				err = errors.New("process exited before listeners became ready")
			}
			return err
		default:
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("listeners not ready within %s", timeout)
		}
		time.Sleep(25 * time.Millisecond)
	}
}

type execMMA2Process struct {
	cmd  *exec.Cmd
	done chan error
}

func startMMA2Process(configPath string) (mma2Process, error) {
	binary := os.Getenv(mma2BinaryEnv)
	if binary == "" {
		var err error
		binary, err = exec.LookPath("mma2")
		if err != nil {
			return nil, fmt.Errorf("MMA2 executable unavailable; set %s: %w", mma2BinaryEnv, err)
		}
	}
	cmd := exec.Command(binary, configPath)
	var stderr bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start %s: %w", binary, err)
	}
	p := &execMMA2Process{cmd: cmd, done: make(chan error, 1)}
	go func() {
		err := cmd.Wait()
		if err != nil && stderr.Len() > 0 {
			err = fmt.Errorf("%w: %s", err, bytes.TrimSpace(stderr.Bytes()))
		}
		p.done <- err
		close(p.done)
	}()
	return p, nil
}

func (p *execMMA2Process) Stop() error {
	if p.cmd.Process == nil {
		return nil
	}
	if err := p.cmd.Process.Signal(os.Interrupt); err != nil && !errors.Is(err, os.ErrProcessDone) {
		return err
	}
	select {
	case <-p.done:
		return nil
	case <-time.After(time.Second):
		if err := p.cmd.Process.Kill(); err != nil && !errors.Is(err, os.ErrProcessDone) {
			return err
		}
		<-p.done
		return nil
	}
}

func (p *execMMA2Process) Done() <-chan error { return p.done }
