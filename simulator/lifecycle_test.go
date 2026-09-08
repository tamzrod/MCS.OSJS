package simulator

import (
	"errors"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

type listenerProcess struct {
	listeners []net.Listener
	done      chan error
	once      sync.Once
}

func (p *listenerProcess) Stop() error {
	p.once.Do(func() {
		for _, listener := range p.listeners {
			_ = listener.Close()
		}
		close(p.done)
	})
	return nil
}

func (p *listenerProcess) Done() <-chan error { return p.done }

func listenerFactory(configPath string) (mma2Process, error) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg EffectiveMMA2Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	p := &listenerProcess{done: make(chan error, 1)}
	for _, address := range effectiveListeners(cfg) {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			_ = p.Stop()
			return nil, err
		}
		p.listeners = append(p.listeners, listener)
	}
	return p, nil
}

func TestMMA2LifecycleActivationReplacementAndRollback(t *testing.T) {
	store := Store{Root: t.TempDir()}
	firstPort := reservePort(t)
	secondPort := reservePort(t)
	first := EffectiveMMA2Config{Listeners: []MMA2Listener{{ID: "first", Listen: firstPort}}}
	if err := store.saveEffective(first); err != nil {
		t.Fatal(err)
	}
	lifecycle := &MMA2Lifecycle{factory: listenerFactory, readyTimeout: time.Second}
	defer lifecycle.Stop()
	if err := lifecycle.Activate(store.EffectiveConfigPath(), first); err != nil {
		t.Fatal(err)
	}
	assertListening(t, firstPort)

	second := EffectiveMMA2Config{Listeners: []MMA2Listener{{ID: "second", Listen: secondPort}}}
	if err := store.saveEffective(second); err != nil {
		t.Fatal(err)
	}
	if err := lifecycle.Activate(store.EffectiveConfigPath(), second); err != nil {
		t.Fatal(err)
	}
	assertListening(t, secondPort)
	assertNotListening(t, firstPort)

	failures := 0
	lifecycle.factory = func(path string) (mma2Process, error) {
		failures++
		if failures == 1 {
			return nil, errors.New("forced start failure")
		}
		return listenerFactory(path)
	}
	third := EffectiveMMA2Config{Listeners: []MMA2Listener{{ID: "third", Listen: reservePort(t)}}}
	if err := store.saveEffective(third); err != nil {
		t.Fatal(err)
	}
	if err := lifecycle.Activate(store.EffectiveConfigPath(), third); err == nil {
		t.Fatal("expected replacement failure")
	}
	assertListening(t, secondPort)
}

func reservePort(t *testing.T) string {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	address := listener.Addr().String()
	_ = listener.Close()
	return address
}

func assertListening(t *testing.T, address string) {
	t.Helper()
	conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)
	if err != nil {
		t.Fatalf("%s not listening: %v", address, err)
	}
	_ = conn.Close()
}

func assertNotListening(t *testing.T, address string) {
	t.Helper()
	conn, err := net.DialTimeout("tcp", address, 50*time.Millisecond)
	if err == nil {
		_ = conn.Close()
		t.Fatalf("%s still listening", address)
	}
}
