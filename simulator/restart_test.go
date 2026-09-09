package simulator

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func acknowledgeNextRestart(store Store) {
	go func() {
		deadline := time.Now().Add(3 * time.Second)
		for time.Now().Before(deadline) {
			if request, found, _ := store.LoadRestartRequest(); found {
				_ = os.WriteFile(store.RestartAckPath(), []byte(request.ConfigSHA256), 0o644)
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
}

func freePort(t *testing.T) uint16 {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	_, portText, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	return uint16(parsePort(portText, t))
}

func parsePort(s string, t *testing.T) int {
	t.Helper()
	var n int
	if _, err := fmt.Sscan(s, &n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestStructuralApplyRequestsRestartAndClearsAfterReadiness(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	_, portText, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	store := Store{Root: t.TempDir()}
	device := validDevice()
	device.Name = "restart-ready-device"
	port := uint16(parsePort(portText, t))
	device.MMA2.Port = port
	applier := newSchedulerApplier(store, Document{}, false)
	defer applier.Stop()
	acknowledgeNextRestart(store)

	if err := applier.ApplyStructural(Document{}, Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatalf("structural apply should succeed against ready listener: %v", err)
	}

	// Readiness confirmed: no restart request may linger after success。
	if _, found, err := store.LoadRestartRequest(); err != nil {
		t.Fatal(err)
	} else if found {
		t.Fatal("restart request must be cleared after confirmed readiness")
	}

	cfg, err := store.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Listeners) != 1 {
		t.Fatalf("composed config missing the simulator reservation: %+v", cfg)
	}
}

func TestStructuralApplyRejectedCommitWritesNoRestartRequest(t *testing.T) {
	store := Store{Root: t.TempDir()}
	first := validDevice()
	first.Name = "dup-a"
	second := validDevice()
	second.Name = "dup-b"
	applier := newSchedulerApplier(store, Document{}, false)
	defer applier.Stop()

	// Same (port,unit_id( twice: compose must reject before any write.

	err := applier.ApplyStructural(Document{}, Document{Devices: []DeviceDefinition{first, second}})
	if err == nil {
		t.Fatal("expected duplicate reservation rejection")
	}
	if _, found, err := store.LoadRestartRequest(); err != nil {
		t.Fatal(err)
	} else if found {
		t.Fatal("rejected config commit must not write a restart request")
	}
}

func TestStructuralApplyRestartTimeoutReportsFailureAndLeavesPendingRequest(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDevice()
	device.Name = "restart-timeout-device"
	device.MMA2.Port = freePort(t)
	applier := newSchedulerApplier(store, Document{}, false)
	applier.restartTimeout = 300 * time.Millisecond
	defer applier.Stop()
	acknowledgeNextRestart(store)

	err := applier.ApplyStructural(Document{}, Document{Devices: []DeviceDefinition{device}})
	if err == nil {
		t.Fatal("expected restart readiness failure")
	}
	if !strings.Contains(err.Error(), "restart not ready") {
		t.Fatalf("failure must surface restart unreadiness truthfully: %v", err)
	}
	// Unconfirmed request stays persisted as truthful evidence of the pending restart。
	if _, found, err := store.LoadRestartRequest(); err != nil {
		t.Fatal(err)
	} else if !found {
		t.Fatal("unconfirmed restart request must remain pending")
	}
}

func TestWaitMMA2Ready(t *testing.T) {
	// A listening appliance port is ready immediately。
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	_, portText, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	port := uint16(parsePort(portText, t))
	if err := WaitMMA2Ready([]uint16{port}, 200*time.Millisecond); err != nil {
		t.Fatalf("live listener must be ready: %v", err)
	}
	ln.Close()

	// A closed port must not claim readiness。
	if err := WaitMMA2Ready([]uint16{freePort(t)}, 200*time.Millisecond); err == nil {
		t.Fatal("closed port must not report ready")
	}
}
