package simulator

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

// PersistenceRBEHandler receives one-byte persistence RBE rule events. The live
// SchedulerApplier implements it via PublishPersistenceEvent.
type PersistenceRBEHandler interface {
	PublishPersistenceEvent(id uint8)
}

// PersistenceRBESubscriber consumes the MMA2 RBE v1 one-byte event stream and
// routes each event to the persistence save handler. Events are ephemeral (there
// is no replay); the subscriber only forwards rule IDs. A zero byte (reserved,
// not a valid rule ID) is ignored. It never writes MMA2 memory and never
// fabricates events.
type PersistenceRBESubscriber struct {
	handler PersistenceRBEHandler
}

// NewPersistenceRBESubscriber builds a subscriber over the given handler.
func NewPersistenceRBESubscriber(handler PersistenceRBEHandler) (*PersistenceRBESubscriber, error) {
	if handler == nil {
		return nil, fmt.Errorf("persistence RBE subscriber requires a handler")
	}
	return &PersistenceRBESubscriber{handler: handler}, nil
}

// Subscribe reads one-byte rule IDs from r until it ends or ctx is cancelled,
// forwarding each non-zero ID to the handler. It returns nil on ctx cancellation
// or clean EOF, and the read error otherwise.
func (s *PersistenceRBESubscriber) Subscribe(ctx context.Context, r io.Reader) error {
	buf := make([]byte, 1)
	for {
		if err := ctx.Err(); err != nil {
			return nil
		}
		n, err := r.Read(buf)
		if n == 1 && buf[0] != 0 {
			s.handler.PublishPersistenceEvent(buf[0])
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// DialAndSubscribe connects to the MMA2 RBE v1 endpoint and forwards its events
// to the handler until ctx is cancelled. A dial failure is returned so the
// caller can decide whether persistence save wiring is required.
func (s *PersistenceRBESubscriber) DialAndSubscribe(ctx context.Context, endpoint string, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	conn, err := net.DialTimeout("tcp", endpoint, timeout)
	if err != nil {
		return fmt.Errorf("persistence RBE dial %s: %w", endpoint, err)
	}
	defer conn.Close()
	go func() {
		<-ctx.Done()
		_ = conn.Close()
	}()
	return s.Subscribe(ctx, conn)
}
