package simulator

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBridgeRoundTrip(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}
	want := Document{Devices: []DeviceDefinition{validDevice(), secondDevice()}}
	ts := httptest.NewServer(NewBridge(store))
	defer ts.Close()
	body := mustJSON(t, want)
	resp := mustPut(t, ts.URL+BridgePath, body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("PUT status %d", resp.StatusCode)
	}
	resp.Body.Close()
	got, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	assertDevicesEqual(t, got.Devices, want.Devices)
	resp, err = http.Get(ts.URL + BridgePath)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var gotDoc Document
	if err := json.Unmarshal(data, &gotDoc); err != nil {
		t.Fatal(err)
	}
	assertDevicesEqual(t, gotDoc.Devices, want.Devices)
}

func TestBridgeRejectsInvalidAtomic(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}
	good := validDevice()
	err := store.SaveOne(good)
	if err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(NewBridge(store))
	defer ts.Close()
	bad := secondDevice()
	bad.MMA2.FC3.Start = 65530
	badDoc := Document{Devices: []DeviceDefinition{good, bad}}
	body := mustJSON(t, badDoc)
	resp := mustPut(t, ts.URL+BridgePath, body)
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("PUT invalid status %d", resp.StatusCode)
	}
	resp.Body.Close()
	got, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Devices) != 1 {
		t.Fatalf("device count %d", len(got.Devices))
	}
	if got.Devices[0] != good {
		t.Fatalf("device changed: got %v", got.Devices[0])
	}
}

func TestApplyingBridgeReturnsSelectedPath(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}
	base := Document{Devices: []DeviceDefinition{validDevice()}}
	if err := store.SaveDocument(base); err != nil {
		t.Fatal(err)
	}
	recorder := &applyRecorder{}
	router := NewApplyRouter(store, recorder, recorder)
	ts := httptest.NewServer(NewApplyingBridge(store, router))
	defer ts.Close()
	edited := cloneDocument(base)
	edited.Devices[0].RandomRuntime.FC1IntervalMS++
	resp := mustPut(t, ts.URL+BridgePath, mustJSON(t, edited))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("PUT status %d", resp.StatusCode)
	}
	var result ApplyResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if result.Path != ApplyTiming || recorder.timing != 1 || recorder.structural != 0 {
		t.Fatalf("apply result=%+v recorder=%+v", result, recorder)
	}
}

func mustJSON(t *testing.T, v interface{}) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func assertDevicesEqual(t *testing.T, got, want []DeviceDefinition) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("device count %d != %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("device %d mismatch: want %v, got %v", i, want[i], got[i])
		}
	}
}

func mustPut(t *testing.T, url string, body string) *http.Response {
	t.Helper()
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return resp
}
