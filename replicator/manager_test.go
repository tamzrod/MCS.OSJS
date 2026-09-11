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
	ackNext := func() <-chan error {
		ch := make(chan error, 1)
		go func() {
			deadline := time.Now().Add(1500 * time.Millisecond)
			for time.Now().Before(deadline) {
				body, readErr := os.ReadFile(requestPath)
				if readErr == nil {
					var req struct {
						ConfigSHA256 string `yaml:"config_sha256"`
					}
					if unmarshalErr := yaml.Unmarshal(body, &req); unmarshalErr != nil {
						ch <- unmarshalErr
						return
					}
					ch <- os.WriteFile(ackPath, []byte(req.ConfigSHA256), 0o644)
					return
				}
				if !os.IsNotExist(readErr) {
					ch <- readErr
					return
				}
				time.Sleep(10 * time.Millisecond)
			}
			ch <- os.ErrDeadlineExceeded
		}()
		return ch
	}

	device := validDeviceDefinition("PLC-runtime")
	device.Endpoint = net.JoinHostPort(sourceHost, portString(sourcePort))
	device.PullBlocks = []PullBlock{
		{Function: 3, Start: 10, Count: 4, ScanRateMS: 40},
		{Function: 4, Start: 100, Count: 2, ScanRateMS: 180},
	}
	device.PullBlock = device.PullBlocks[0]
	device.Destination = DestinationSelection{Port: destinationPort, UnitID: 9}
	ackDone := ackNext()
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
	if len(doc.Devices) != 1 || len(doc.Devices[0].PullBlocks) != 2 || doc.Devices[0].Destination.Status != "OWNED" || doc.Devices[0].Destination.Owner != ProducerReplicator {
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
		if len(status.Blocks) == 2 && status.Blocks[0].Cycles >= 4 && status.Blocks[1].Cycles >= 1 {
			if !status.Running || status.Source != "ERROR" || status.LastError == "" || status.LastPoll == "" {
				t.Fatalf("status = %#v, want running ERROR with last error/poll", status)
			}
			if status.Blocks[0].Cycles <= status.Blocks[1].Cycles {
				t.Fatalf("independent cadence not visible: fast=%d slow=%d", status.Blocks[0].Cycles, status.Blocks[1].Cycles)
			}
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("runtimes did not complete independent cycles before timeout: %#v", status)
		}
		time.Sleep(20 * time.Millisecond)
	}

	// Save & Apply always restarts MMA2 even when only one block scan rate changes.
	edited := cloneDocument(doc)
	edited.Devices[0].PullBlocks[0].ScanRateMS = 90
	edited.Devices[0].PullBlock = edited.Devices[0].PullBlocks[0]
	ackDone = ackNext()
	_, structural, err = manager.Apply(edited)
	if err != nil {
		t.Fatalf("non-structural Apply: %v", err)
	}
	if structural {
		t.Fatal("scan-rate-only edit must not be structural")
	}
	if ackErr := <-ackDone; ackErr != nil {
		t.Fatalf("second restart acknowledgement helper: %v", ackErr)
	}
	if _, err := os.Stat(requestPath); !os.IsNotExist(err) {
		t.Fatalf("restart request should be cleared after second apply, stat err=%v", err)
	}
}
