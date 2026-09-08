package simulator

import (
	"bufio"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	RuntimeProtocolVersion = 1
	RuntimeSocketRelPath   = "run/modbus-simulator.sock"
	maxRuntimeMessage      = 1 << 20
)

type runtimeApplier interface {
	Apply(Document) (ApplyResult, error)
}

type runtimeStatusReader interface {
	RuntimeStatus(string) (DeviceRuntimeStatus, error)
}

type RuntimeRequest struct {
	Version   int             `json:"version"`
	RequestID string          `json:"request_id"`
	Operation string          `json:"operation"`
	Payload   json.RawMessage `json:"payload"`
}

type RuntimeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RuntimeResponse struct {
	Version   int           `json:"version"`
	RequestID string        `json:"request_id"`
	OK        bool          `json:"ok"`
	Result    any           `json:"result,omitempty"`
	Error     *RuntimeError `json:"error,omitempty"`
}

type ApplyRuntimeResult struct {
	ApplyResult
	CompletedAt time.Time `json:"completed_at"`
}

type RuntimeService struct {
	store   Store
	apply   runtimeApplier
	status  runtimeStatusReader
	mutate  sync.Mutex
	cacheMu sync.Mutex
	cache   map[string]RuntimeResponse
	lastMu  sync.RWMutex
	last    *ApplyRuntimeResult
}

func NewRuntimeService(store Store, apply runtimeApplier, status runtimeStatusReader) *RuntimeService {
	return &RuntimeService{store: store, apply: apply, status: status, cache: make(map[string]RuntimeResponse)}
}

func NewLiveRuntimeService(store Store) (*RuntimeService, *SchedulerApplier, error) {
	router, scheduler, err := NewRuntimeApplyRouter(store)
	if err != nil {
		return nil, nil, err
	}
	return NewRuntimeService(store, router, scheduler), scheduler, nil
}

func (s *RuntimeService) Handle(req RuntimeRequest) RuntimeResponse {
	base := RuntimeResponse{Version: RuntimeProtocolVersion, RequestID: req.RequestID}
	if req.Version != RuntimeProtocolVersion || strings.TrimSpace(req.RequestID) == "" {
		return runtimeFailure(base, "INVALID_REQUEST", "version 1 and a non-empty request_id are required")
	}
	if cached, ok := s.cached(req.RequestID); ok {
		return cached
	}

	var response RuntimeResponse
	switch req.Operation {
	case "load":
		doc, err := s.store.Load()
		if err != nil {
			response = runtimeFailure(base, classifyRuntimeError(err), err.Error())
		} else {
			response = base
			response.OK = true
			response.Result = struct {
				Document  Document            `json:"document"`
				LastApply *ApplyRuntimeResult `json:"last_apply,omitempty"`
			}{doc, s.lastApply()}
		}
	case "apply":
		var payload struct {
			Document Document `json:"document"`
		}
		if err := json.Unmarshal(req.Payload, &payload); err != nil {
			response = runtimeFailure(base, "INVALID_REQUEST", "apply payload requires a document")
			break
		}
		s.mutate.Lock()
		result, err := s.apply.Apply(payload.Document)
		if err == nil {
			completed := ApplyRuntimeResult{ApplyResult: result, CompletedAt: time.Now().UTC()}
			s.lastMu.Lock()
			s.last = &completed
			s.lastMu.Unlock()
			response = base
			response.OK = true
			response.Result = completed
		} else {
			response = runtimeFailure(base, classifyRuntimeError(err), err.Error())
		}
		s.mutate.Unlock()
	case "status":
		var payload struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(req.Payload, &payload); err != nil || strings.TrimSpace(payload.Name) == "" {
			response = runtimeFailure(base, "INVALID_REQUEST", "status payload requires a device name")
			break
		}
		status, err := s.status.RuntimeStatus(payload.Name)
		if err != nil {
			response = runtimeFailure(base, classifyRuntimeError(err), err.Error())
		} else {
			response = base
			response.OK = true
			response.Result = struct {
				Status    DeviceRuntimeStatus `json:"status"`
				LastApply *ApplyRuntimeResult `json:"last_apply,omitempty"`
			}{status, s.lastApply()}
		}
	default:
		response = runtimeFailure(base, "INVALID_REQUEST", "operation must be load, apply, or status")
	}
	s.remember(req.RequestID, response)
	return response
}

func (s *RuntimeService) cached(id string) (RuntimeResponse, bool) {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	result, ok := s.cache[id]
	return result, ok
}

func (s *RuntimeService) remember(id string, response RuntimeResponse) {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	if len(s.cache) >= 256 {
		s.cache = make(map[string]RuntimeResponse)
	}
	s.cache[id] = response
}

func (s *RuntimeService) lastApply() *ApplyRuntimeResult {
	s.lastMu.RLock()
	defer s.lastMu.RUnlock()
	if s.last == nil {
		return nil
	}
	copy := *s.last
	return &copy
}

func runtimeFailure(base RuntimeResponse, code, message string) RuntimeResponse {
	base.Error = &RuntimeError{Code: code, Message: message}
	return base
}

func classifyRuntimeError(err error) string {
	message := strings.ToLower(err.Error())
	switch {
	case errors.Is(err, ErrReservationOwnedByOther):
		return "RESERVATION_CONFLICT"
	case strings.Contains(message, "restart") && strings.Contains(message, "not ready"):
		return "MMA2_NOT_READY"
	case strings.Contains(message, "restart"):
		return "MMA2_RESTART_FAILED"
	case strings.Contains(message, "raw ingest"):
		return "RAW_INGEST_FAILED"
	case strings.Contains(message, "required"), strings.Contains(message, "must"), strings.Contains(message, "invalid"), strings.Contains(message, "duplicate"):
		return "VALIDATION_FAILED"
	default:
		return "INTERNAL"
	}
}

func RuntimeSocketPath(root string) string {
	return filepath.Join(root, RuntimeSocketRelPath)
}

func ServeRuntime(ctx context.Context, socketPath string, service *RuntimeService) error {
	if err := os.MkdirAll(filepath.Dir(socketPath), 0o750); err != nil {
		return err
	}
	if _, err := os.Stat(socketPath); err == nil {
		conn, dialErr := net.DialTimeout("unix", socketPath, 150*time.Millisecond)
		if dialErr == nil {
			_ = conn.Close()
			return fmt.Errorf("runtime socket already has a live owner: %s", socketPath)
		}
		if err := os.Remove(socketPath); err != nil {
			return fmt.Errorf("remove stale runtime socket: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return err
	}
	defer listener.Close()
	defer os.Remove(socketPath)
	if err := os.Chmod(socketPath, 0o660); err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = listener.Close()
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			return err
		}
		go handleRuntimeConn(conn, service)
	}
}

func handleRuntimeConn(conn net.Conn, service *RuntimeService) {
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(10 * time.Second))
	reader := bufio.NewReader(conn)
	var size uint32
	if err := binary.Read(reader, binary.BigEndian, &size); err != nil || size == 0 || size > maxRuntimeMessage {
		return
	}
	body := make([]byte, size)
	if _, err := io.ReadFull(reader, body); err != nil {
		return
	}
	var request RuntimeRequest
	if err := json.Unmarshal(body, &request); err != nil {
		writeRuntimeResponse(conn, runtimeFailure(RuntimeResponse{Version: RuntimeProtocolVersion}, "INVALID_REQUEST", "request is not valid JSON"))
		return
	}
	writeRuntimeResponse(conn, service.Handle(request))
}

func writeRuntimeResponse(writer io.Writer, response RuntimeResponse) {
	body, err := json.Marshal(response)
	if err != nil {
		return
	}
	_ = binary.Write(writer, binary.BigEndian, uint32(len(body)))
	_, _ = writer.Write(body)
}
