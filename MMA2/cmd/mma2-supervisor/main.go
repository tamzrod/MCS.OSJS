package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"gopkg.in/yaml.v3"
	"mma2/internal/restartwatch"
)

const pollInterval = 50 * time.Millisecond

type engine struct {
	binary string
	config string
	cmd    *exec.Cmd
	done   chan error
}

func (e *engine) start() error {
	e.cmd = exec.Command(e.binary, e.config)
	e.cmd.Stdout = os.Stdout
	e.cmd.Stderr = os.Stderr
	e.done = make(chan error, 1)
	if err := e.cmd.Start(); err != nil {
		return err
	}
	go func() { e.done <- e.cmd.Wait() }()
	return nil
}

func (e *engine) stop() error {
	if e.cmd == nil || e.cmd.Process == nil {
		return nil
	}
	_ = e.cmd.Process.Signal(syscall.SIGTERM)
	select {
	case err := <-e.done:
		e.cmd = nil
		return err
	case <-time.After(2 * time.Second):
		_ = e.cmd.Process.Kill()
		err := <-e.done
		e.cmd = nil
		return err
	}
}

func main() {
	binary, configPath, requestPath, ackPath := resolveRuntimePaths()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	e := &engine{binary: binary, config: configPath}
	initialRequest, err := os.ReadFile(requestPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("read initial restart request: %v", err)
	}
	var tracker restartwatch.Tracker
	tracker.Seed(initialRequest)
	if err := e.start(); err != nil {
		log.Fatalf("start MMA2: %v", err)
	}
	if len(initialRequest) > 0 {
		if err := acknowledge(initialRequest, ackPath); err != nil {
			log.Fatalf("acknowledge initial restart request: %v", err)
		}
	}
	log.Printf("MMA2 appliance started with %s", e.config)

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			_ = e.stop()
			return
		case err := <-e.done:
			log.Fatalf("MMA2 exited unexpectedly: %v", err)
		case <-ticker.C:
			request, readErr := os.ReadFile(requestPath)
			if errors.Is(readErr, os.ErrNotExist) {
				tracker.Observe(nil)
				continue
			}
			if readErr != nil {
				log.Printf("read restart request: %v", readErr)
				continue
			}
			if !tracker.Observe(request) {
				continue
			}
			sum := sha256.Sum256(request)
			log.Printf("restart request consumed fingerprint=%s", hex.EncodeToString(sum[:8]))
			_ = e.stop()
			if err := e.start(); err != nil {
				log.Fatalf("restart MMA2: %v", err)
			}
			if err := acknowledge(request, ackPath); err != nil {
				log.Fatalf("acknowledge restart request: %v", err)
			}
		}
	}
}

func resolveRuntimePaths() (binary, configPath, requestPath, ackPath string) {
	if len(os.Args) == 5 {
		return os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	}
	if len(os.Args) != 1 {
		log.Fatal("usage: mma2-supervisor <mma2-binary> <config.yaml> <restart-request.yaml> <restart-ack>")
	}

	root := os.Getenv("OSJS_DATA_DIR")
	if root == "" && runtime.GOOS == "windows" {
		if programData := os.Getenv("ProgramData"); programData != "" {
			root = filepath.Join(programData, "MCS Modbus Toolkit", "runtime")
		}
	}
	if root == "" {
		log.Fatal("OSJS_DATA_DIR is required when mma2-supervisor runs without command-line arguments")
	}

	// Keep child/runtime instrumentation aligned with the resolved shared root.
	_ = os.Setenv("OSJS_DATA_DIR", root)
	_ = os.Setenv("MCS_DATA_ROOT", root)

	supervisorPath, err := os.Executable()
	if err != nil {
		log.Fatalf("resolve supervisor executable: %v", err)
	}
	binDir := filepath.Dir(supervisorPath)
	mma2Dir := filepath.Join(root, "config", "mma2")
	return filepath.Join(binDir, "mma2.exe"),
		filepath.Join(mma2Dir, "config.yaml"),
		filepath.Join(mma2Dir, "restart-request.yaml"),
		filepath.Join(mma2Dir, "restart-ack")
}

func acknowledge(request []byte, path string) error {
	var parsed struct {
		ConfigSHA256 string `yaml:"config_sha256"`
	}
	if err := yaml.Unmarshal(request, &parsed); err != nil {
		return err
	}
	if parsed.ConfigSHA256 == "" {
		return errors.New("restart request has no config_sha256")
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(parsed.ConfigSHA256), 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
