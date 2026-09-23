package simulator

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v3"
	"mma2/pkg/configvalidate"
)

type SharedMMASettings struct {
	Settings            map[string]interface{} `json:"settings"`
	Revision            string                 `json:"revision"`
	Committed           bool                   `json:"committed"`
	RestartAcknowledged bool                   `json:"restart_acknowledged"`
	Message             string                 `json:"message,omitempty"`
}

type sharedMMAError struct {
	code    string
	message string
}

func (err *sharedMMAError) Error() string      { return err.message }
func sharedFailure(code, message string) error { return &sharedMMAError{code: code, message: message} }
func sharedKey(key string) bool                { return key == "rbe" || key == "access_events" || key == "debug" }

func sharedSnapshot(data []byte, config map[string]interface{}) SharedMMASettings {
	sum := sha256.Sum256(data)
	result := SharedMMASettings{Settings: map[string]interface{}{}, Revision: hex.EncodeToString(sum[:])}
	for key, value := range config {
		if sharedKey(key) {
			result.Settings[key] = value
		}
	}
	return result
}

func (s Store) readSharedMMA() (SharedMMASettings, map[string]interface{}, error) {
	data, err := os.ReadFile(s.EffectiveConfigPath())
	if err != nil && !os.IsNotExist(err) {
		return SharedMMASettings{}, nil, err
	}
	config := map[string]interface{}{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return SharedMMASettings{}, nil, err
	}
	if config == nil {
		config = map[string]interface{}{}
	}
	return sharedSnapshot(data, config), config, nil
}

func (s Store) loadSharedMMA() (SharedMMASettings, error) {
	var result SharedMMASettings
	err := s.withWriterLock(func() error {
		var err error
		result, _, err = s.readSharedMMA()
		return err
	})
	return result, err
}

func (s Store) applySharedMMA(revision string, settings map[string]interface{}, restart func(RestartRequest) error) (SharedMMASettings, error) {
	var result SharedMMASettings
	if len(revision) != 64 || settings == nil {
		return result, sharedFailure("INVALID_REQUEST", "revision and settings object are required")
	}
	for key, value := range settings {
		if !sharedKey(key) {
			return result, sharedFailure("INVALID_REQUEST", "unsupported shared setting: "+key)
		}
		if value == nil {
			continue
		}
		if key == "debug" {
			if _, ok := value.(bool); !ok {
				return result, sharedFailure("VALIDATION_FAILED", "debug must be boolean")
			}
		} else if _, ok := value.(map[string]interface{}); !ok {
			return result, sharedFailure("VALIDATION_FAILED", key+" must be an object or null")
		}
	}
	err := s.withWriterLock(func() error {
		current, config, err := s.readSharedMMA()
		if err != nil {
			return err
		}
		result = current
		if revision != current.Revision {
			return sharedFailure("REVISION_CONFLICT", "MMA settings changed; reload before saving")
		}
		for key, value := range settings {
			if value == nil {
				delete(config, key)
			} else {
				config[key] = value
			}
		}
		data, err := yaml.Marshal(config)
		if err != nil {
			return err
		}
		if err := configvalidate.YAML(data); err != nil {
			return sharedFailure("VALIDATION_FAILED", err.Error())
		}
		if err := s.composer().WriteFileAtomic(s.EffectiveConfigPath(), data); err != nil {
			return err
		}
		result = sharedSnapshot(data, config)
		result.Committed = true
		request := RestartRequest{RequestedAt: time.Now().UTC(), Reason: "Shared MMA settings commit", ConfigSHA256: result.Revision}
		if err := s.writeRestartRequestLocked(request); err != nil {
			return sharedFailure("MMA2_RESTART_FAILED", fmt.Sprintf("configuration committed; restart request failed; recovery required: %v", err))
		}
		if err := restart(request); err != nil {
			return sharedFailure("MMA2_RESTART_FAILED", fmt.Sprintf("configuration committed but restart not acknowledged; recovery required: %v", err))
		}
		result.RestartAcknowledged = true
		if err := s.clearRestartRequestLocked(); err != nil {
			return sharedFailure("MMA2_RESTART_FAILED", "restart acknowledged but request cleanup failed: "+err.Error())
		}
		result.Message = "MMA settings saved; restart acknowledged. Runtime readiness is not verified; check status and diagnostics."
		return nil
	})
	return result, err
}

func (s *RuntimeService) handleSharedMMA(req RuntimeRequest, base RuntimeResponse) (response RuntimeResponse) {
	s.mutate.Lock()
	defer s.mutate.Unlock()
	if cached, ok := s.cached(req.RequestID); ok {
		return cached
	}
	defer func() { s.remember(req.RequestID, response) }()
	var result SharedMMASettings
	var err error
	if req.Operation == "mma-load" {
		result, err = s.store.loadSharedMMA()
	} else {
		var payload struct {
			Revision string                 `json:"revision"`
			Settings map[string]interface{} `json:"settings"`
		}
		decoder := json.NewDecoder(bytes.NewReader(req.Payload))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&payload); err != nil {
			return runtimeFailure(base, "INVALID_REQUEST", "shared apply requires revision and settings")
		}
		if err := decoder.Decode(new(interface{})); err != io.EOF {
			return runtimeFailure(base, "INVALID_REQUEST", "invalid trailing payload")
		}
		result, err = s.store.applySharedMMA(payload.Revision, payload.Settings, func(request RestartRequest) error {
			return s.store.WaitRestartAcknowledged(request.ConfigSHA256, DefaultRestartReadyTimeout)
		})
	}
	base.Result = result
	if err != nil {
		code := "INTERNAL"
		if typed, ok := err.(*sharedMMAError); ok {
			code = typed.code
		}
		return runtimeFailure(base, code, err.Error())
	}
	base.OK = true
	return base
}
