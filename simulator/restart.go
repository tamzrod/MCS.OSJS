package simulator

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	restartRequestFile = "restart-request.yaml"
	restartAckFile = "restart-ack"
	DefaultRestartReadyTimeout = 20 * time.Second
	restartReadyPoll = 25 * time.Millisecond
)

// RestartRequest is the only lifecycle command the Simulator issues to the
// independently owned MMA2 service. A request never starts or kills MMA2.
type RestartRequest struct {
	RequestedAt time.Time `yaml:"requested_at"`
	Reason string `yaml:"reason"`
	ConfigSHA256 string `yaml:"config_sha256"`
	Ports []uint16 `yaml:"ports"`
}

func (s Store) RestartRequestPath() string { return filepath.Join(s.MMA2ConfigDir(), restartRequestFile) }
func (s Store) RestartAckPath() string { return filepath.Join(s.MMA2ConfigDir(), restartAckFile) }

func RestartRequestForConfig(now time.Time, cfg EffectiveMMA2Config, ports []uint16) RestartRequest {
	b, _ := yaml.Marshal(&cfg)
	sum := sha256.Sum256(b)
	return RestartRequest{
		RequestedAt: now.UTC(),
		Reason: "Simulator shared MMA2 configuration commit",
		ConfigSHA256: hex.EncodeToString(sum[:]),
		Ports: append([]uint16(nil), ports...),
	}
}

// Public standalone restart commands acquire the lock. Structural apply
// already owns the shared transaction and uses the private locked variant.
func (s Store) WriteRestartRequest(req RestartRequest) error {
	return s.withWriterLock(func() error { return s.writeRestartRequestLocked(req) })
}

func (s Store) writeRestartRequestLocked(req RestartRequest) error {
	if err := os.Remove(s.RestartAckPath()); err != nil && !os.IsNotExist(err) { return err }
	b, err := yaml.Marshal(&req)
	if err != nil { return err }
	return s.replaceMMA2File(s.RestartRequestPath(), b)
}

func (s Store) WaitRestartAcknowledged(configSHA256 string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for {
		ack, err := os.ReadFile(s.RestartAckPath())
		if err == nil && string(ack) == configSHA256 { return nil }
		if err != nil && !os.IsNotExist(err) { return fmt.Errorf("read MMA2 restart acknowledgement: %w", err) }
		if time.Now().After(deadline) { return fmt.Errorf("mma2 restart request was not acknowledged within timeout") }
		time.Sleep(restartReadyPoll)
	}
}

func (s Store) LoadRestartRequest() (RestartRequest, bool, error) {
	b, err := os.ReadFile(s.RestartRequestPath())
	if err != nil {
		if os.IsNotExist(err) { return RestartRequest{}, false, nil }
		return RestartRequest{}, false, err
	}
	var req RestartRequest
	if err := yaml.Unmarshal(b, &req); err != nil { return RestartRequest{}, false, fmt.Errorf("parse mma2 restart request %s: %w", s.RestartRequestPath(), err) }
	return req, true, nil
}

func (s Store) ClearRestartRequest() error {
	return s.withWriterLock(func() error { return s.clearRestartRequestLocked() })
}

func (s Store) clearRestartRequestLocked() error {
	if err := os.Remove(s.RestartRequestPath()); err != nil && !os.IsNotExist(err) { return err }
	if err := os.Remove(s.RestartAckPath()); err != nil && !os.IsNotExist(err) { return err }
	return nil
}

func composedSimulatorPorts(doc Document) []uint16 {
	seen := make(map[uint16]bool)
	var ports []uint16
	for _, def := range doc.Devices {
		if !def.Enabled || !hasMMA2Areas(def.MMA2) || seen[def.MMA2.Port] { continue }
		seen[def.MMA2.Port] = true
		ports = append(ports, def.MMA2.Port)
	}
	return ports
}

func WaitMMA2Ready(ports []uint16, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for _, port := range ports {
		if err := waitMMA2PortReady(port, deadline); err != nil { return err }
	}
	return nil
}

func waitMMA2PortReady(port uint16, deadline time.Time) error {
	addr := net.JoinHostPort("127.0.0.1", fmt.Sprint(port))
	for {
		conn, err := net.DialTimeout("tcp", addr, time.Until(deadline))
		if err == nil { _ = conn.Close(); return nil }
		if time.Now().After(deadline) { return fmt.Errorf("mma2 restart not ready within timeout on %s: %w", addr, err) }
		time.Sleep(restartReadyPoll)
	}
}
