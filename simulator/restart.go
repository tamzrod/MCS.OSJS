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
	// restartRequestFile is the machine-readable restart request artifact stored
	// beside the shared MMA2 configuration after a successful structural commit。
	restartRequestFile = "restart-request.yaml"

	// DefaultRestartReadyTimeout bounds SIM-014's readiness wait after a
	// restart request。
	DefaultRestartReadyTimeout = 20 * time.Second

	restartReadyPoll = 25 * time.Millisecond
)

// RestartRequest is the single MMA2 lifecycle/control request the Simulator may
// issue: exactly one RESTART for the just-committed shared configuration. The
// Simulator never starts, stops, spawns, kills, or replaces MMA2 itself。
type RestartRequest struct {
	RequestedAt  time.Time `yaml:"requested_at"`
	Reason       string    `yaml:"reason"`
	ConfigSHA256 string    `yaml:"config_sha256"`
	Ports        []uint16  `yaml:"ports"`
}

func (s Store) RestartRequestPath() string {
	return filepath.Join(s.MMA2ConfigDir(), restartRequestFile)
}

// RestartRequestForConfig builds a restart request for the composed shared config,
// fingerprinting the exact committed content so a consumer can tell whether the
// request corresponds to the currently committed config。
func RestartRequestForConfig(now time.Time, cfg EffectiveMMA2Config, ports []uint16) RestartRequest {
	b, _ := yaml.Marshal(&cfg)
	sum := sha256.Sum256(b)
	out := RestartRequest{
		RequestedAt:  now.UTC(),
		Reason:       "Simulator shared MMA2 configuration commit",
		ConfigSHA256: hex.EncodeToString(sum[:]),
		Ports:        append([]uint16(nil), ports...),
	}
	return out
}

// WriteRestartRequest atomically persists the restart request beside the shared
// MMA2 config artifacts。
func (s Store) WriteRestartRequest(req RestartRequest) error {
	b, err := yaml.Marshal(&req)
	if err != nil {
		return err
	}
	return s.replaceMMA2File(s.RestartRequestPath(), b)
}

// LoadRestartRequest reads the persisted restart request;a missing file means
// no restart is pending。
func (s Store) LoadRestartRequest() (RestartRequest, bool, error) {
	b, err := os.ReadFile(s.RestartRequestPath())
	if err != nil {
		if os.IsNotExist(err) {
			return RestartRequest{}, false, nil
		}
		return RestartRequest{}, false, err
	}
	var req RestartRequest
	if err := yaml.Unmarshal(b, &req); err != nil {
		return RestartRequest{}, false, fmt.Errorf("parse mma2 restart request %s: %w", s.RestartRequestPath(), err)
	}
	return req, true, nil
}

// ClearRestartRequest removes a pending restart request after readiness has been
// confirmed, leaving no stale request that could mislead a later boot restore。
func (s Store) ClearRestartRequest() error {
	if err := os.Remove(s.RestartRequestPath()); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// composedSimulatorPorts lists the MMA2 listener ports the simulator composed for
// enabled devices with memory areas, deduplicated,in the order the shared config
// serves them。
func composedSimulatorPorts(doc Document) []uint16 {
	seen := make(map[uint16]bool)
	var ports []uint16
	for _, def := range doc.Devices {
		if !def.Enabled || !hasMMA2Areas(def.MMA2) {
			continue
		}
		if seen[def.MMA2.Port] {
			continue
		}
		seen[def.MMA2.Port] = true
		ports = append(ports, def.MMA2.Port)
	}
	return ports
}

// WaitMMA2Ready waits for every composed simulator MMA2 listener to accept a
// TCP dial, returning a truthful error when the independently managed appliance
// has not returned ready within timeout。
func WaitMMA2Ready(ports []uint16, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for _, port := range ports {
		if err := waitMMA2PortReady(port, deadline); err != nil {
			return err
		}
	}
	return nil
}

func waitMMA2PortReady(port uint16, deadline time.Time) error {
	addr := net.JoinHostPort("127.0.0.1", fmt.Sprint(port))
	for {
		conn, err := net.DialTimeout("tcp", addr, time.Until(deadline))
		if err == nil {
			_ = conn.Close()
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("mma2 restart not ready within timeout on %s: %w", addr, err)
		}
		time.Sleep(restartReadyPoll)
	}
}
