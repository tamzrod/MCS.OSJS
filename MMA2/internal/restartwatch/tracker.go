package restartwatch

import "crypto/sha256"

// Tracker recognizes each on-disk restart request once while it remains
// present. Removing the request resets the tracker for a later request.
type Tracker struct {
	seen    [sha256.Size]byte
	present bool
}

func (t *Tracker) Observe(request []byte) bool {
	if len(request) == 0 {
		t.present = false
		return false
	}
	sum := sha256.Sum256(request)
	if t.present && sum == t.seen {
		return false
	}
	t.seen = sum
	t.present = true
	return true
}

func (t *Tracker) Seed(request []byte) {
	if len(request) == 0 {
		return
	}
	t.seen = sha256.Sum256(request)
	t.present = true
}
