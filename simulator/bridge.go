package simulator

import (
	"encoding/json"
	"net/http"
)

// BridgePath is the loopback HTTP bridge root for simulator-owned document
// persistence (SIM-005). Browser form edits round-trip through the SIM-001 Go
// model via this handler pair: GET loads the persisted Document, PUT atomically
// replaces it through Store.SaveDocument. The bridge never touches effective MMA2
// configuration (backend activation is SIM-006 scope(.
const BridgePath = "/api/devices"

// Bridge serves the simulator-owned device document over JSON on an internal
// loopback listener. It is a transport only: validation lives in ValidateDevice.
type Bridge struct {
	store   Store
	applier *ApplyRouter
}

func NewBridge(store Store) *Bridge {
	return &Bridge{store: store}
}

func NewApplyingBridge(store Store, applier *ApplyRouter) *Bridge {
	return &Bridge{store: store, applier: applier}
}

func (b Bridge) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == BridgePath:
		b.get(w, r)
	case r.Method == http.MethodPut && r.URL.Path == BridgePath:
		b.put(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (b Bridge) get(w http.ResponseWriter, r *http.Request) {
	doc, err := b.store.Load()
	if err != nil {
		http.Error(w, "load failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, doc)
}

func (b Bridge) put(w http.ResponseWriter, r *http.Request) {
	var doc Document
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&doc); err != nil {
		http.Error(w, "invalid document JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if b.applier != nil {
		result, err := b.applier.Apply(doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		writeJSON(w, http.StatusOK, result)
		return
	}
	if err := b.store.SaveDocument(doc); err != nil {
		// SaveDocument errors come from validation or the atomic replace; both
		// leave prior persisted bytes unchanged and indicate the submitted document
		// was not applied, so surface the reason as unprocessable.
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	saved, err := b.store.Load()
	if err != nil {
		http.Error(w, "reload failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, saved)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		// Response already committed; nothing more we can do.
		_ = err
	}
}
