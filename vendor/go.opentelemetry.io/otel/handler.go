// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otel // import "go.opentelemetry.io/otel"

import (
<<<<<<< HEAD
	"log"
	"os"
	"sync"
	"sync/atomic"
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	"log"
	"os"
	"sync"
=======
	"go.opentelemetry.io/otel/internal/global"
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
)

var (
<<<<<<< HEAD
	// globalErrorHandler provides an ErrorHandler that can be used
	// throughout an OpenTelemetry instrumented project. When a user
	// specified ErrorHandler is registered (`SetErrorHandler`) all calls to
	// `Handle` and will be delegated to the registered ErrorHandler.
	globalErrorHandler = &loggingErrorHandler{
		l: log.New(os.Stderr, "", log.LstdFlags),
	}

	// delegateErrorHandlerOnce ensures that a user provided ErrorHandler is
	// only ever registered once.
	delegateErrorHandlerOnce sync.Once

	// Comiple time check that loggingErrorHandler implements ErrorHandler.
	_ ErrorHandler = (*loggingErrorHandler)(nil)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	// globalErrorHandler provides an ErrorHandler that can be used
	// throughout an OpenTelemetry instrumented project. When a user
	// specified ErrorHandler is registered (`SetErrorHandler`) all calls to
	// `Handle` and will be delegated to the registered ErrorHandler.
	globalErrorHandler = defaultErrorHandler()

	// Compile-time check that delegator implements ErrorHandler.
	_ ErrorHandler = (*delegator)(nil)
	// Compile-time check that errLogger implements ErrorHandler.
	_ ErrorHandler = (*errLogger)(nil)
=======
	// Compile-time check global.ErrDelegator implements ErrorHandler.
	_ ErrorHandler = (*global.ErrDelegator)(nil)
	// Compile-time check global.ErrLogger implements ErrorHandler.
	_ ErrorHandler = (*global.ErrLogger)(nil)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
)

<<<<<<< HEAD
// loggingErrorHandler logs all errors to STDERR.
type loggingErrorHandler struct {
	delegate atomic.Value

	l *log.Logger
}

// setDelegate sets the ErrorHandler delegate if one is not already set.
func (h *loggingErrorHandler) setDelegate(d ErrorHandler) {
	if h.delegate.Load() != nil {
		// Delegate already registered
		return
	}
	h.delegate.Store(d)
}

// Handle implements ErrorHandler.
func (h *loggingErrorHandler) Handle(err error) {
	if d := h.delegate.Load(); d != nil {
		d.(ErrorHandler).Handle(err)
		return
	}
	h.l.Print(err)
}

// GetErrorHandler returns the global ErrorHandler instance. If no ErrorHandler
// instance has been set (`SetErrorHandler`), the default ErrorHandler which
// logs errors to STDERR is returned.
func GetErrorHandler() ErrorHandler {
	return globalErrorHandler
}
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
type delegator struct {
	lock *sync.RWMutex
	eh   ErrorHandler
}

func (d *delegator) Handle(err error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	d.eh.Handle(err)
}

// setDelegate sets the ErrorHandler delegate.
func (d *delegator) setDelegate(eh ErrorHandler) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.eh = eh
}

func defaultErrorHandler() *delegator {
	return &delegator{
		lock: &sync.RWMutex{},
		eh:   &errLogger{l: log.New(os.Stderr, "", log.LstdFlags)},
	}
}

// errLogger logs errors if no delegate is set, otherwise they are delegated.
type errLogger struct {
	l *log.Logger
}

// Handle logs err if no delegate is set, otherwise it is delegated.
func (h *errLogger) Handle(err error) {
	h.l.Print(err)
}

// GetErrorHandler returns the global ErrorHandler instance.
//
// The default ErrorHandler instance returned will log all errors to STDERR
// until an override ErrorHandler is set with SetErrorHandler. All
// ErrorHandler returned prior to this will automatically forward errors to
// the set instance instead of logging.
//
// Subsequent calls to SetErrorHandler after the first will not forward errors
// to the new ErrorHandler for prior returned instances.
func GetErrorHandler() ErrorHandler {
	return globalErrorHandler
}
=======
// GetErrorHandler returns the global ErrorHandler instance.
//
// The default ErrorHandler instance returned will log all errors to STDERR
// until an override ErrorHandler is set with SetErrorHandler. All
// ErrorHandler returned prior to this will automatically forward errors to
// the set instance instead of logging.
//
// Subsequent calls to SetErrorHandler after the first will not forward errors
// to the new ErrorHandler for prior returned instances.
func GetErrorHandler() ErrorHandler { return global.GetErrorHandler() }
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

<<<<<<< HEAD
// SetErrorHandler sets the global ErrorHandler to be h.
func SetErrorHandler(h ErrorHandler) {
	delegateErrorHandlerOnce.Do(func() {
		current := GetErrorHandler()
		if current == h {
			return
		}
		if internalHandler, ok := current.(*loggingErrorHandler); ok {
			internalHandler.setDelegate(h)
		}
	})
}
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// SetErrorHandler sets the global ErrorHandler to h.
//
// The first time this is called all ErrorHandler previously returned from
// GetErrorHandler will send errors to h instead of the default logging
// ErrorHandler. Subsequent calls will set the global ErrorHandler, but not
// delegate errors to h.
func SetErrorHandler(h ErrorHandler) {
	globalErrorHandler.setDelegate(h)
}
=======
// SetErrorHandler sets the global ErrorHandler to h.
//
// The first time this is called all ErrorHandler previously returned from
// GetErrorHandler will send errors to h instead of the default logging
// ErrorHandler. Subsequent calls will set the global ErrorHandler, but not
// delegate errors to h.
func SetErrorHandler(h ErrorHandler) { global.SetErrorHandler(h) }
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

<<<<<<< HEAD
// Handle is a convience function for ErrorHandler().Handle(err)
func Handle(err error) {
	GetErrorHandler().Handle(err)
}
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// Handle is a convenience function for ErrorHandler().Handle(err).
func Handle(err error) {
	GetErrorHandler().Handle(err)
}
=======
// Handle is a convenience function for ErrorHandler().Handle(err).
func Handle(err error) { global.Handle(err) }
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
