package replicator

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"gopkg.in/yaml.v3"
)

const (
	restartRequestFile = "restart-request.yaml"
	restartAckFile     = "restart-ack"
	restartPoll        = 25 * time.Millisecond
	DefaultApplyTimeout = 20 * time.Second
)

type restartRequest struct {
	RequestedAt  time.Time `yaml:"requested_at"`
	Reason       string    `yaml:"reason"`
	ConfigSHA256 string    `yaml:"config_sha256"`
	Ports        []uint16  `yaml:"ports"`
}

func (s Store) requestMMA2Restart(cfg mma2composer.EffectiveConfig, ports []uint16, timeout time.Duration) error {
	composer := mma2composer.New(s.Root, ProducerReplicator)
	configBytes, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	sum := sha256.Sum256(configBytes)
	req := restartRequest{
		RequestedAt:  time.Now().UTC(),
		Reason:       "Replicator shared MMA2 configuration commit",
		ConfigSHA256: hex.EncodeToString(sum[:]),
		Ports:        append([]uint16(nil), ports...),
	}
	configDir := composer.ConfigDir()
	requestPath := filepath.Join(configDir, restartRequestFile)
	ackPath := filepath.Join(configDir, restartAckFile)
	if err := os.Remove(ackPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	body, err := yaml.Marshal(&req)
	if err != nil {
		return err
	}
	if err := composer.WriteFileAtomic(requestPath, body); err != nil {
		return fmt.Errorf("write MMA2 restart request: %w", err)
	}
	if err := waitRestartAck(ackPath, req.ConfigSHA256, timeout); err != nil {
		return err
	}
	if err := waitPortsReady(ports, timeout); err != nil {
		return err
	}
	if err := os.Remove(requestPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.Remove(ackPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func waitRestartAck(path, fingerprint string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for {
		ack, err := os.ReadFile(path)
		if err == nil && string(ack) == fingerprint {
			return nil
		}
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("read MMA2 restart acknowledgement: %w", err)
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("mma2 restart request was not acknowledged within timeout")
		}
		time.Sleep(restartPoll)
	}
}

func waitPortsReady(ports []uint16, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for _, port := range ports {
		address := net.JoinHostPort("127.0.0.1", fmt.Sprint(port))
		for {
			conn, err := net.DialTimeout("tcp", address, 250*time.Millisecond)
			if err == nil {
				_ = conn.Close()
				break
			}
			if time.Now().After(deadline) {
				return fmt.Errorf("mma2 not ready within timeout on %s: %w", address, err)
			}
			time.Sleep(restartPoll)
		}
	}
	return nil
}
