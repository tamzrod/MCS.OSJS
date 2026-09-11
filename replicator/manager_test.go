package replicator

import (
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"gopkg.in/yaml.v3"
)

func TestRuntimeManagerApplyLifecycleAndStatus(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}

	destinationListener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer destinationListener.Close()
	_, destinationPort := splitTestAddress(t, destinationListener.Addr().String())

	// Reserve a source port and release it so the runtime gets a deterministic
	// connection-refused cycle after apply, proving truthful ERROR status.
	sourceListener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	sourceHost, sourcePort := splitTestAddress(t, sourceListener.Addr().String())
	_ = sourceListener.Close()

	manager := NewRuntimeManager(store)
	manager.timeout = 2 * time.Second
	defer manager.Stop()

	composer := mma2composer.New(root, ProducerReplicator)
	requestPath := filepath.Join(composer.ConfigDir(), restartRequestFile)
	ackPath := filepath.Join(composer.ConfigDir(), restartAckFile)
	ackDone := make(chan error, 1)
	go func() {
		deadline := time.Now().Add(1500 * time.Millisecond)
		for time.Now().Before(deadline) {
			body, readErr := os.ReadFile(requestPath)
			if readErr == nil {
				var req struct {
					ConfigSHA256 string `yaml:"config_sha256"`
				}
				if unmarshalErr := yaml.Unmarshal(body, &req); unmarshalErr != nil {
					ackDone <- unmarshalErr
					return
				}
				ackDone <- os.WriteFile(ackPath, []byte(req.ConfigSHA256), 0o644)
				return
			}
			if !os.IsNotExist(readErr) {
				ackDone <- readErr
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
		ackDone <- os.ErrDeadlineExceeded
	}()

	device := validDeviceDefinition("PLC-runtime")
	device.Endpoint = net.JoinHostPort(sourceHost, portString(sourcePort))
	device.Destination = DestinationSelection{Port: destinationPort, UnitID: 9}
	doc, structural, err := manager.Apply(Document{Devices: []DeviceDefinition{device}})
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if !structural {
		t.Fatal("first apply must be structural")
	}
	if ackErr := <-ackDone; ackErr != nil {
		t.Fatalf("restart acknowledgement helper: %v", ackErr)
	}
	if len(doc.Devices) != 1 || doc.Devices[0].Destination.Status != "OWNED" || doc.Devices[0].Destination.Owner != ProducerReplicator {
		t.Fatalf("resolved document = %#v", doc)
	}
	if _, err := os.Stat(requestPath); !os.IsNotExist(err) {
		t.Fatalf("restart request should be cleared after readiness, stat err=%v", err)
	}

	deadline := time.Now().Add(1500 * time.Millisecond)
	for {
		status, statusErr := manager.Status(device.Name)
		if statusErr != nil {
			t.Fatal(statusErr)
		}
		if status.Cycles > 0 {
			if !status.Running || status.Source != "ERROR" || status.LastError == "" || status.LastPoll == "" {
				t.Fatalf("status = %#v, want running ERROR with last error/poll", status)
			}
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("runtime did not complete a cycle before timeout: %#v", status)
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Scan-rate-only edit is non-structural: it restarts the Replicator poll loop
	// but must not request an MMA2 restart.
	edited := doc
	edited.Devices[0].ScanRateMS = 500
	_, structural, err = manager.Apply(edited)
	if err != nil {
		t.Fatalf("non-structural Apply: %v", err)
	}
	if structural {
		t.Fatal("scan-rate-only edit must not be structural")
	}
	if _, err := os.Stat(requestPath); !os.IsNotExist(err) {
		t.Fatalf("non-structural apply wrote restart request, stat err=%v", err)
	}
}
