package simulator

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"
)

func sharedFixture(t *testing.T) Store {
	t.Helper()
	store := Store{Root: t.TempDir()}
	if err := store.SaveAndCompose(validDevice()); err != nil {
		t.Fatal(err)
	}
	cfg, err := store.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Extra == nil {
		cfg.Extra = map[string]interface{}{}
	}
	cfg.Extra["custom_root"] = "preserve"
	if err := store.composer().SaveEffective(cfg); err != nil {
		t.Fatal(err)
	}
	return store
}

func sharedFile(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestSharedMMAApplyPreservesForeignConfiguration(t *testing.T) {
	store := sharedFixture(t)
	before, err := store.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	owners := sharedFile(t, store.OwnershipPath())
	devices := sharedFile(t, store.DevicesPath())
	initial, err := store.loadSharedMMA()
	if err != nil {
		t.Fatal(err)
	}
	settings := map[string]interface{}{"debug": true, "rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:39001"}}}
	calls := 0
	result, err := store.applySharedMMA(initial.Revision, settings, func(request RestartRequest) error {
		calls++
		current, loadErr := store.loadEffective()
		if loadErr != nil {
			return loadErr
		}
		if !reflect.DeepEqual(current.Listeners, before.Listeners) {
			t.Fatal("listeners changed")
		}
		pending, exists, loadErr := store.LoadRestartRequest()
		if loadErr != nil || !exists || pending.ConfigSHA256 != request.ConfigSHA256 {
			t.Fatal("missing or unmatched restart request")
		}
		data := sharedFile(t, store.EffectiveConfigPath())
		if sharedSnapshot(data, nil).Revision != request.ConfigSHA256 {
			t.Fatal("acknowledgment revision does not match actual saved bytes")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Committed || !result.RestartAcknowledged || calls != 1 || result.Revision == initial.Revision {
		t.Fatalf("bad result: %+v", result)
	}
	after, err := store.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if after.Extra["custom_root"] != "preserve" {
		t.Fatal("unknown root extension lost")
	}
	if !bytes.Equal(owners, sharedFile(t, store.OwnershipPath())) || !bytes.Equal(devices, sharedFile(t, store.DevicesPath())) {
		t.Fatal("foreign artifacts changed")
	}
	if _, exists, err := store.LoadRestartRequest(); err != nil || exists {
		t.Fatal("acknowledged request not cleared")
	}
	loaded, err := store.loadSharedMMA()
	if err != nil || loaded.Revision != result.Revision {
		t.Fatalf("reload: %+v %v", loaded, err)
	}
}

func TestSharedMMARejectsConflictAndInvalidCandidatesWithoutWrites(t *testing.T) {
	store := sharedFixture(t)
	initial, err := store.loadSharedMMA()
	if err != nil {
		t.Fatal(err)
	}
	before := sharedFile(t, store.EffectiveConfigPath())
	for _, check := range []struct {
		name, revision, code string
		patch                map[string]interface{}
	}{
		{"conflict", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "REVISION_CONFLICT", map[string]interface{}{"debug": true}},
		{"foreign field", initial.Revision, "INVALID_REQUEST", map[string]interface{}{"listeners": []interface{}{}}},
		{"debug type", initial.Revision, "VALIDATION_FAILED", map[string]interface{}{"debug": "yes"}},
		{"invalid RBE", initial.Revision, "VALIDATION_FAILED", map[string]interface{}{"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "bad-address"}}}},
	} {
		t.Run(check.name, func(t *testing.T) {
			result, err := store.applySharedMMA(check.revision, check.patch, func(RestartRequest) error { t.Fatal("restart on rejected candidate"); return nil })
			var failure *sharedMMAError
			if !errors.As(err, &failure) || failure.code != check.code || result.Committed {
				t.Fatalf("result=%+v error=%v", result, err)
			}
			if !bytes.Equal(before, sharedFile(t, store.EffectiveConfigPath())) {
				t.Fatal("rejection changed disk")
			}
		})
	}
}

func TestSharedMMACommittedButUnacknowledged(t *testing.T) {
	store := sharedFixture(t)
	initial, err := store.loadSharedMMA()
	if err != nil {
		t.Fatal(err)
	}
	result, err := store.applySharedMMA(initial.Revision, map[string]interface{}{"debug": true}, func(RestartRequest) error { return errors.New("test timeout") })
	var failure *sharedMMAError
	if !errors.As(err, &failure) || failure.code != "MMA2_RESTART_FAILED" || !result.Committed || result.RestartAcknowledged {
		t.Fatalf("result=%+v error=%v", result, err)
	}
	loaded, err := store.loadSharedMMA()
	if err != nil || loaded.Revision != result.Revision || loaded.Settings["debug"] != true {
		t.Fatal("committed config was hidden or rolled back")
	}
	if _, exists, err := store.LoadRestartRequest(); err != nil || !exists {
		t.Fatal("pending request not preserved")
	}
}

func TestSharedMMANullRemovalAndRuleValidation(t *testing.T) {
	store := sharedFixture(t)
	initial, err := store.loadSharedMMA()
	if err != nil {
		t.Fatal(err)
	}
	withOutput, err := store.applySharedMMA(initial.Revision, map[string]interface{}{"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:39001"}}}, func(RestartRequest) error { return nil })
	if err != nil {
		t.Fatal(err)
	}
	removed, err := store.applySharedMMA(withOutput.Revision, map[string]interface{}{"rbe": nil}, func(RestartRequest) error { return nil })
	if err != nil {
		t.Fatal(err)
	}
	if _, exists := removed.Settings["rbe"]; exists {
		t.Fatal("explicit null did not remove output")
	}
	withOutput, err = store.applySharedMMA(removed.Revision, map[string]interface{}{"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:39001"}}}, func(RestartRequest) error { return nil })
	if err != nil {
		t.Fatal(err)
	}
	device := validDevice()
	device.MMA2.RBE = map[string]interface{}{"coils": []interface{}{map[string]interface{}{"id": 1, "name": "watch", "start": 0, "count": 1}}}
	if err := store.SaveAndCompose(device); err != nil {
		t.Fatal(err)
	}
	current, err := store.loadSharedMMA()
	if err != nil {
		t.Fatal(err)
	}
	result, err := store.applySharedMMA(current.Revision, map[string]interface{}{"rbe": nil}, func(RestartRequest) error { t.Fatal("invalid candidate restarted"); return nil })
	var failure *sharedMMAError
	if !errors.As(err, &failure) || failure.code != "VALIDATION_FAILED" || result.Committed {
		t.Fatalf("result=%+v error=%v", result, err)
	}
}

func TestSharedMMARuntimeContract(t *testing.T) {
	store := sharedFixture(t)
	service := NewRuntimeService(store, &runtimeApplyRecorder{}, runtimeStatusFixture{})
	response := service.Handle(RuntimeRequest{Version: 1, RequestID: "shared-load", Operation: "mma-load", Payload: json.RawMessage(`{}`)})
	if !response.OK {
		t.Fatalf("load failed: %+v", response.Error)
	}
	value, ok := response.Result.(SharedMMASettings)
	if !ok || len(value.Revision) != 64 {
		t.Fatal("missing snapshot")
	}
	response = service.Handle(RuntimeRequest{Version: 1, RequestID: "bad-shared", Operation: "mma-apply", Payload: json.RawMessage(`{"revision":"bad","settings":{},"listeners":[]}`)})
	if response.OK || response.Error.Code != "INVALID_REQUEST" {
		t.Fatal("accepted unsupported top-level mutation")
	}
}

func TestSharedMMARuntimeAcknowledgmentFile(t *testing.T) {
	store := sharedFixture(t)
	service := NewRuntimeService(store, &runtimeApplyRecorder{}, runtimeStatusFixture{})
	initial, err := store.loadSharedMMA()
	if err != nil {
		t.Fatal(err)
	}
	finished := make(chan error, 1)
	go func() {
		deadline := time.Now().Add(3 * time.Second)
		for time.Now().Before(deadline) {
			request, exists, err := store.LoadRestartRequest()
			if err != nil {
				finished <- err
				return
			}
			if exists {
				finished <- os.WriteFile(store.RestartAckPath(), []byte(request.ConfigSHA256), 0600)
				return
			}
			time.Sleep(time.Millisecond)
		}
		finished <- errors.New("test restart request not received")
	}()
	payload, err := json.Marshal(map[string]interface{}{"revision": initial.Revision, "settings": map[string]interface{}{"debug": true}})
	if err != nil {
		t.Fatal(err)
	}
	request := RuntimeRequest{Version: 1, RequestID: "shared-ack", Operation: "mma-apply", Payload: payload}
	responses := make(chan RuntimeResponse, 2)
	for index := 0; index < 2; index++ {
		go func() { responses <- service.Handle(request) }()
	}
	response := <-responses
	retry := <-responses
	if !retry.OK {
		t.Fatalf("concurrent retry failed: %+v", retry.Error)
	}
	if err := <-finished; err != nil {
		t.Fatal(err)
	}
	if !response.OK {
		t.Fatalf("shared apply failed: %+v", response.Error)
	}
	result := response.Result.(SharedMMASettings)
	if !result.Committed || !result.RestartAcknowledged {
		t.Fatalf("unacknowledged result: %+v", result)
	}
	if cached := service.Handle(RuntimeRequest{Version: 1, RequestID: "shared-ack", Operation: "mma-apply", Payload: payload}); !cached.OK {
		t.Fatal("retry replayed mutation rather than cached acknowledgment")
	}
	if _, exists, err := store.LoadRestartRequest(); err != nil || exists {
		t.Fatal("request retained after acknowledgment")
	}
}
