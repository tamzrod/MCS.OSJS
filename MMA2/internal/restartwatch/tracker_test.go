package restartwatch

import "testing"

func TestTrackerConsumesEachRequestExactlyOnce(t *testing.T) {
	var tracker Tracker
	first := []byte("requested_at: first")
	if !tracker.Observe(first) {
		t.Fatal("first request was not consumed")
	}
	if tracker.Observe(first) {
		t.Fatal("unchanged request was consumed twice")
	}
	if !tracker.Observe([]byte("requested_at: second")) {
		t.Fatal("changed request was not consumed")
	}
	if tracker.Observe(nil) {
		t.Fatal("request removal must not trigger a restart")
	}
	if !tracker.Observe(first) {
		t.Fatal("request after removal was not consumed")
	}
}

func TestTrackerSeedPreventsRedundantStartupRestart(t *testing.T) {
	request := []byte("pending request at container startup")
	var tracker Tracker
	tracker.Seed(request)
	if tracker.Observe(request) {
		t.Fatal("startup request caused a redundant second launch")
	}
}
