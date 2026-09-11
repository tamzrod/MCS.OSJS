package replicator

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"
)

const RuntimeProtocolVersion = 1

type RuntimeRequest struct {
	Version   int             `json:"version"`
	RequestID string          `json:"request_id"`
	Operation string          `json:"operation"`
	Payload   json.RawMessage `json:"payload,omitempty"`
}

type RuntimeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RuntimeResponse struct {
	Version   int           `json:"version"`
	RequestID string        `json:"request_id"`
	OK        bool          `json:"ok"`
	Result    interface{}   `json:"result,omitempty"`
	Error     *RuntimeError `json:"error,omitempty"`
}

type ApplyResponse struct {
	Document   Document `json:"document"`
	Structural bool     `json:"structural"`
	Message    string   `json:"message"`
	CompletedAt string  `json:"completed_at"`
}

type LoadResponse struct {
	Document   Document              `json:"document"`
	Suggestion DestinationSuggestion `json:"suggestion"`
}

type destinationQuery struct {
	Inspect bool   `json:"inspect"`
	Port    uint16 `json:"port"`
	UnitID  uint16 `json:"unit_id"`
}

type statusQuery struct {
	Name string `json:"name"`
}

type applyQuery struct {
	Document Document `json:"document"`
}

func RuntimeSocketPath(root string) string {
	return filepath.Join(root, "run", "modbus-replicator.sock")
}

func HandleRuntimeRequest(manager *RuntimeManager, request RuntimeRequest) RuntimeResponse {
	response := RuntimeResponse{Version: RuntimeProtocolVersion, RequestID: request.RequestID}
	if request.Version != RuntimeProtocolVersion {
		response.Error = &RuntimeError{Code: "INVALID_VERSION", Message: "unsupported Replicator runtime protocol version"}
		return response
	}
	fail := func(code string, err error) RuntimeResponse {
		response.Error = &RuntimeError{Code: code, Message: err.Error()}
		return response
	}

	switch request.Operation {
	case "load":
		suggestion, err := manager.store.SuggestDestination()
		if err != nil {
			return fail("LOAD_FAILED", err)
		}
		response.OK = true
		response.Result = LoadResponse{Document: manager.Document(), Suggestion: suggestion}
	case "suggest":
		var query destinationQuery
		if len(request.Payload) > 0 {
			if err := json.Unmarshal(request.Payload, &query); err != nil {
				return fail("INVALID_REQUEST", err)
			}
		}
		var suggestion DestinationSuggestion
		var err error
		if query.Inspect {
			suggestion, err = manager.store.InspectDestination(query.Port, query.UnitID)
		} else {
			suggestion, err = manager.store.SuggestDestination()
		}
		if err != nil {
			return fail("SUGGEST_FAILED", err)
		}
		response.OK = true
		response.Result = suggestion
	case "status":
		var query statusQuery
		if err := json.Unmarshal(request.Payload, &query); err != nil {
			return fail("INVALID_REQUEST", err)
		}
		status, err := manager.Status(query.Name)
		if err != nil {
			return fail("STATUS_FAILED", err)
		}
		response.OK = true
		response.Result = status
	case "apply":
		var query applyQuery
		if err := json.Unmarshal(request.Payload, &query); err != nil {
			return fail("INVALID_REQUEST", err)
		}
		doc, structural, err := manager.Apply(query.Document)
		if err != nil {
			return fail("APPLY_FAILED", err)
		}
		message := "Replicator settings applied without an MMA2 structural restart."
		if structural {
			message = "Replicator destination structure applied through the MMA2 restart path."
		}
		response.OK = true
		response.Result = ApplyResponse{Document: doc, Structural: structural, Message: message, CompletedAt: time.Now().UTC().Format(time.RFC3339Nano)}
	default:
		return fail("INVALID_REQUEST", fmt.Errorf("unsupported operation %q", request.Operation))
	}
	return response
}
