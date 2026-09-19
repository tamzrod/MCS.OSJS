package replicator

import (
	"encoding/json"
	"io"
	"net"
	"testing"
	"time"
)

func TestCommsCycleAndStatus(t *testing.T) {
	for _, destinationOK := range []bool{true, false} {
		sourceAddress, stopSource := startTestModbusServer(t, 3, 12, []uint16{42})
		host, port := splitTestAddress(t, sourceAddress)
		destinationAddress, received, stopDestination := startRawCaptureServer(t, 1)
		_, destinationPort := splitTestAddress(t, destinationAddress)
		if !destinationOK {
			stopDestination()
		}
		cfg := Config{Source: SourceConfig{Host: host, Port: port, UnitID: 7, Function: 3, Start: 12, Count: 1},
			Destination: DestinationConfig{ListenerPort: destinationPort, UnitID: 9, Area: "fc3", Start: 12, Count: 1}}
		runtime := newManagedRuntime(cfg)
		runtime.setRunning(true)
		runtime.runCycle()
		stopSource()
		if destinationOK {
			<-received
			stopDestination()
		}
		device := validDeviceDefinition("LED")
		device.PullBlocks = []PullBlock{{Function: 3, Start: 12, Count: 1, ScanRateMS: 100}}
		manager := NewRuntimeManager(Store{})
		manager.document = Document{Devices: []DeviceDefinition{device}}
		manager.runtimes["LED"] = []*managedRuntime{runtime}
		status, err := manager.Status("LED")
		if err != nil {
			t.Fatal(err)
		}
		for _, layer := range []string{"network", "tcp", "modbus"} {
			if status.Comms[layer] != "OK" {
				t.Fatalf("%s = %s", layer, status.Comms[layer])
			}
		}
		want := "ERROR"
		if destinationOK {
			want = "OK"
		}
		if status.Comms["mma2"] != want {
			t.Fatalf("MMA2 = %s, want %s", status.Comms["mma2"], want)
		}
		block := status.Blocks[0]
		if block.Function != 3 || block.Start != 12 || block.Count != 1 || block.TCP.ActivityAt == "" {
			t.Fatalf("incomplete block: %+v", block)
		}
		if (block.MMA2.LastSuccessAt != "") != destinationOK {
			t.Fatal("unacknowledged destination activity")
		}
		body, err := json.Marshal(status)
		if err != nil {
			t.Fatal(err)
		}
		var wire map[string]any
		if err := json.Unmarshal(body, &wire); err != nil {
			t.Fatal(err)
		}
		wireBlock := wire["blocks"].([]any)[0].(map[string]any)
		if wireBlock["tcp"].(map[string]any)["state"] != "OK" {
			t.Fatalf("renderer contract: %s", body)
		}
	}
}

func TestCommsSourceRefusalClearsDownstream(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	host, port := splitTestAddress(t, listener.Addr().String())
	listener.Close()
	var comms CycleComms
	_, err = readSourceObserved(SourceConfig{Host: host, Port: port, UnitID: 1, Function: 3, Count: 1}, 100*time.Millisecond, &comms)
	if err == nil || comms.TCP.State != "ERROR" || comms.Network.State != "UNKNOWN" || comms.Modbus.State == "OK" || comms.MMA2.State == "OK" {
		t.Fatalf("unexpected evidence: %+v, %v", comms, err)
	}
}

func TestCommsExceptionAndMalformedResponse(t *testing.T) {
	for _, exception := range []bool{true, false} {
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		host, port := splitTestAddress(t, listener.Addr().String())
		done := make(chan struct{})
		go func() {
			defer close(done)
			connection, acceptErr := listener.Accept()
			if acceptErr != nil {
				return
			}
			defer connection.Close()
			connection.SetDeadline(time.Now().Add(time.Second))
			request := make([]byte, 12)
			if _, readErr := io.ReadFull(connection, request); readErr != nil {
				return
			}
			function := byte(0x83)
			if !exception {
				function = 4
			}
			connection.Write([]byte{0, 1, 0, 0, 0, 3, 1, function, 2})
		}()
		var comms CycleComms
		_, err = readSourceObserved(SourceConfig{Host: host, Port: port, UnitID: 1, Function: 3, Count: 1}, time.Second, &comms)
		listener.Close()
		<-done
		want := "ERROR"
		if exception {
			want = "WARNING"
		}
		if err == nil || comms.TCP.State != "OK" || comms.Modbus.State != want {
			t.Fatalf("unexpected evidence: %+v, %v", comms, err)
		}
		if exception && (comms.Modbus.ExceptionCode == nil || *comms.Modbus.ExceptionCode != 2) {
			t.Fatal("missing exception evidence")
		}
	}
}

func TestCommsAggregation(t *testing.T) {
	for _, testCase := range []struct {
		states []string
		want   string
	}{
		{nil, "UNKNOWN"}, {[]string{"OK", "OK"}, "OK"}, {[]string{"OK", ""}, "UNKNOWN"},
		{[]string{"OK", "WARNING"}, "WARNING"}, {[]string{"ERROR", "OK"}, "ERROR"},
	} {
		observations := []CommsObservation{}
		for _, state := range testCase.states {
			observations = append(observations, CommsObservation{State: state})
		}
		if got := aggregateComms(observations); got != testCase.want {
			t.Fatalf("%v: %s", testCase.states, got)
		}
	}
}
