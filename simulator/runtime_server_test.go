package simulator

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

type runtimeApplyRecorder struct {
	mu    sync.Mutex
	calls int
}

func (r *runtimeApplyRecorder) Apply(doc Document) (ApplyResult, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls++
	return ApplyResult{Document: doc, Path: ApplyTiming, Message: "applied"}, nil
}

type runtimeStatusFixture struct{}

func (runtimeStatusFixture) RuntimeStatus(name string) (DeviceRuntimeStatus, error) {
	if name == "missing" {
		return DeviceRuntimeStatus{}, fmt.Errorf("device %q not found", name)
	}
	return DeviceRuntimeStatus{Name: name, Device: "ERROR", MMA2: "RUNNING", RawIngest: "ERROR", RawError: "raw ingest rejected frame: code 33", TotalPoints: 1, FC: map[string]FCRuntimeStatus{}}, nil
}

func TestRuntimeServiceStatusContractCarriesOperatorStatesAndDiagnostic(t *testing.T) {
	service := NewRuntimeService(Store{Root: t.TempDir()}, &runtimeApplyRecorder{}, runtimeStatusFixture{})
	payload, _ := json.Marshal(map[string]string{"name": "runtime-device"})
	response := service.Handle(RuntimeRequest{Version: 1, RequestID: "status-contract", Operation: "status", Payload: payload})
	if !response.OK {
		t.Fatalf("status failed: %+v", response.Error)
	}

	wire, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	var decoded struct {
		Result struct {
			Status struct {
				MMA2       string `json:"mma2_status"`
				Simulator  string `json:"device_status"`
				Diagnostic string `json:"raw_ingest_error"`
			} `json:"status"`
		} `json:"result"`
	}
	if err := json.Unmarshal(wire, &decoded); err != nil {
		t.Fatal(err)
	}
	status := decoded.Result.Status
	if status.MMA2 != "RUNNING" || status.Simulator != "ERROR" || status.Diagnostic != "raw ingest rejected frame: code 33" {
		t.Fatalf("status contract changed SIM-021A truth: %+v", status)
	}
}

func TestRuntimeServiceStatusRequestErrorsRemainStable(t *testing.T) {
	service := NewRuntimeService(Store{Root: t.TempDir()}, &runtimeApplyRecorder{}, runtimeStatusFixture{})
	tests := []struct {
		name    string
		payload json.RawMessage
		code    string
		message string
	}{
		{"missing name", json.RawMessage(`{}`), "INVALID_REQUEST", "status payload requires a device name"},
		{"invalid payload", json.RawMessage(`{"name":`), "INVALID_REQUEST", "status payload requires a device name"},
		{"unknown device", json.RawMessage(`{"name":"missing"}`), "INTERNAL", `device "missing" not found`},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := service.Handle(RuntimeRequest{Version: 1, RequestID: fmt.Sprintf("status-error-%d", i), Operation: "status", Payload: test.payload})
			if response.OK || response.Error == nil || response.Error.Code != test.code || response.Error.Message != test.message {
				t.Fatalf("response = %+v, want %s: %s", response, test.code, test.message)
			}
		})
	}
}

func TestRuntimeServiceLoadApplyStatusAndIdempotency(t *testing.T) {
	store := Store{Root: t.TempDir()}
	initial := Document{Devices: []DeviceDefinition{testRuntimeDevice("runtime-device")}}
	if err := store.SaveDocument(initial); err != nil {
		t.Fatal(err)
	}
	recorder := &runtimeApplyRecorder{}
	service := NewRuntimeService(store, recorder, runtimeStatusFixture{})

	load := service.Handle(RuntimeRequest{Version: 1, RequestID: "load-1", Operation: "load", Payload: json.RawMessage(`{}`)})
	if !load.OK {
		t.Fatalf("load failed: %+v", load.Error)
	}
	payload, _ := json.Marshal(map[string]Document{"document": initial})
	request := RuntimeRequest{Version: 1, RequestID: "apply-1", Operation: "apply", Payload: payload}
	if response := service.Handle(request); !response.OK {
		t.Fatalf("apply failed: %+v", response.Error)
	}
	if response := service.Handle(request); !response.OK {
		t.Fatalf("cached apply failed: %+v", response.Error)
	}
	recorder.mu.Lock()
	calls := recorder.calls
	recorder.mu.Unlock()
	if calls != 1 {
		t.Fatalf("duplicate request executed %d applies, want 1", calls)
	}
	statusPayload, _ := json.Marshal(map[string]string{"name": "runtime-device"})
	status := service.Handle(RuntimeRequest{Version: 1, RequestID: "status-1", Operation: "status", Payload: statusPayload})
	if !status.OK {
		t.Fatalf("status failed: %+v", status.Error)
	}
}

func TestRuntimeUnixSocketFramingAndSingleOwner(t *testing.T) {
	root := t.TempDir()
	socketPath := filepath.Join(root, "runtime.sock")
	service := NewRuntimeService(Store{Root: root}, &runtimeApplyRecorder{}, runtimeStatusFixture{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() { done <- ServeRuntime(ctx, socketPath, service) }()
	waitForRuntimeSocket(t, socketPath)

	request := RuntimeRequest{Version: 1, RequestID: "wire-load", Operation: "load", Payload: json.RawMessage(`{}`)}
	response := callRuntimeSocket(t, socketPath, request)
	if !response.OK || response.RequestID != request.RequestID {
		t.Fatalf("unexpected response: %+v", response)
	}
	if err := ServeRuntime(context.Background(), socketPath, service); err == nil {
		t.Fatal("second runtime owner unexpectedly acquired live socket")
	}
	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("runtime server did not stop")
	}
}

func TestLiveRuntimeContractAppliesTimingRejectsConflictAndReportsUnavailable(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	port := uint16(listener.Addr().(*net.TCPAddr).Port)
	root := t.TempDir()
	store := Store{Root: root}
	service, scheduler, err := NewLiveRuntimeService(store)
	if err != nil {
		t.Fatal(err)
	}
	defer scheduler.Stop()
	scheduler.restartTimeout = 100 * time.Millisecond

	device := testRuntimeDevice("live-runtime")
	device.MMA2.Port = port
	device.RandomRuntime.FC3IntervalMS = 60000
	structural := runtimeApplyRequest(t, service, "structural", Document{Devices: []DeviceDefinition{device}})
	if !structural.OK {
		t.Fatalf("structural apply failed: %+v", structural.Error)
	}
	if _, pending, err := store.LoadRestartRequest(); err != nil || pending {
		t.Fatalf("restart request was not cleared after readiness: pending=%v err=%v", pending, err)
	}

	device.RandomRuntime.FC3IntervalMS = 30000
	timing := runtimeApplyRequest(t, service, "timing", Document{Devices: []DeviceDefinition{device}})
	if !timing.OK {
		t.Fatalf("timing apply failed: %+v", timing.Error)
	}
	if _, pending, _ := store.LoadRestartRequest(); pending {
		t.Fatal("timing-only apply wrote a restart request")
	}

	duplicate := device
	duplicate.Name = "duplicate"
	rejected := runtimeApplyRequest(t, service, "duplicate", Document{Devices: []DeviceDefinition{device, duplicate}})
	if rejected.OK || rejected.Error == nil || rejected.Error.Code != "VALIDATION_FAILED" {
		t.Fatalf("duplicate reservation was not rejected truthfully: %+v", rejected)
	}

	unavailable := device
	unavailable.MMA2.Port = unusedTCPPort(t)
	failure := runtimeApplyRequest(t, service, "unavailable", Document{Devices: []DeviceDefinition{unavailable}})
	if failure.OK || failure.Error == nil || failure.Error.Code != "MMA2_NOT_READY" {
		t.Fatalf("unavailable MMA2 was not reported truthfully: %+v", failure)
	}
	statusPayload, _ := json.Marshal(map[string]string{"name": device.Name})
	status := service.Handle(RuntimeRequest{Version: 1, RequestID: "after-failure", Operation: "status", Payload: statusPayload})
	if !status.OK {
		t.Fatalf("prior runtime status disappeared after rejected apply: %+v", status.Error)
	}
}

func runtimeApplyRequest(t *testing.T, service *RuntimeService, id string, doc Document) RuntimeResponse {
	t.Helper()
	payload, err := json.Marshal(map[string]Document{"document": doc})
	if err != nil {
		t.Fatal(err)
	}
	return service.Handle(RuntimeRequest{Version: 1, RequestID: id, Operation: "apply", Payload: payload})
}

func unusedTCPPort(t *testing.T) uint16 {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := uint16(listener.Addr().(*net.TCPAddr).Port)
	_ = listener.Close()
	return port
}

func testRuntimeDevice(name string) DeviceDefinition {
	return DeviceDefinition{
		Name: name, Enabled: true,
		MMA2:          MMA2Params{Port: 5020, UnitID: 1, FC3: Area{Count: 1}},
		RandomRuntime: RandomRuntimeParams{FC3IntervalMS: 1000},
	}
}

func waitForRuntimeSocket(t *testing.T, socketPath string) {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("unix", socketPath, 20*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("runtime socket did not become ready")
}

func callRuntimeSocket(t *testing.T, socketPath string, request RuntimeRequest) RuntimeResponse {
	t.Helper()
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	body, _ := json.Marshal(request)
	if err := binary.Write(conn, binary.BigEndian, uint32(len(body))); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Write(body); err != nil {
		t.Fatal(err)
	}
	var size uint32
	if err := binary.Read(conn, binary.BigEndian, &size); err != nil {
		t.Fatal(err)
	}
	responseBody := make([]byte, size)
	if _, err := io.ReadFull(conn, responseBody); err != nil {
		t.Fatal(err)
	}
	var response RuntimeResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		t.Fatal(err)
	}
	return response
}
