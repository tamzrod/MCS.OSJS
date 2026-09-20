package mma2composer

import (
	"errors"
	"time"
)

// DefaultWriterLockTimeout bounds waiting for another process's MMA2 writer.
const DefaultWriterLockTimeout = 5 * time.Second

// ErrWriterLockTimeout means no shared transaction lock was acquired and the
// caller must return an error without beginning any configuration mutation.
var ErrWriterLockTimeout = errors.New("mma2 shared writer lock timed out")

// WithWriterLock executes action while holding the single shared writer lock
// rooted at <OSJS_DATA_DIR>/config/mma2/.writer.lock on Linux. Every writer must
// acquire ONCE at its outermost load/compose/validate/write/restart/ack entry;
// nested composer helpers must NOT reacquire this non-reentrant lock. A lock
// only serializes cooperating processes sharing the same data root; it does not
// authorize writing, replace revision checks, or make multiple YAML writes one
// atomic filesystem operation. Never delete or replace the lock file itself.
//
// Linux uses a kernel-held advisory lock released on process exit. Windows
// deliberately preserves the existing standalone behavior without a new lock;
// unsupported platforms fail closed. Caller-supplied action errors propagate.
