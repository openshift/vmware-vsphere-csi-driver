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

package trace // import "go.opentelemetry.io/otel/sdk/trace"

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	"go.opentelemetry.io/otel/trace"
=======
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
	"go.opentelemetry.io/otel/trace/noop"
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
)

const (
	defaultTracerName = "go.opentelemetry.io/otel/sdk/tracer"
)

// TODO (MrAlias): unify this API option design:
// https://github.com/open-telemetry/opentelemetry-go/issues/536

// TracerProviderConfig
type TracerProviderConfig struct {
	processors []SpanProcessor

	// sampler is the default sampler used when creating new spans.
	sampler Sampler

	// idGenerator is used to generate all Span and Trace IDs when needed.
	idGenerator IDGenerator

	// spanLimits defines the attribute, event, and link limits for spans.
	spanLimits SpanLimits

	// resource contains attributes representing an entity that produces telemetry.
	resource *resource.Resource
}

type TracerProviderOption func(*TracerProviderConfig)

type TracerProvider struct {
	embedded.TracerProvider

	mu             sync.Mutex
<<<<<<< HEAD
	namedTracer    map[instrumentation.Library]*tracer
	spanProcessors atomic.Value
	sampler        Sampler
	idGenerator    IDGenerator
	spanLimits     SpanLimits
	resource       *resource.Resource
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	namedTracer    map[instrumentation.Scope]*tracer
	spanProcessors atomic.Value

	// These fields are not protected by the lock mu. They are assumed to be
	// immutable after creation of the TracerProvider.
	sampler     Sampler
	idGenerator IDGenerator
	spanLimits  SpanLimits
	resource    *resource.Resource
=======
	namedTracer    map[instrumentation.Scope]*tracer
	spanProcessors atomic.Pointer[spanProcessorStates]

	isShutdown atomic.Bool

	// These fields are not protected by the lock mu. They are assumed to be
	// immutable after creation of the TracerProvider.
	sampler     Sampler
	idGenerator IDGenerator
	spanLimits  SpanLimits
	resource    *resource.Resource
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

var _ trace.TracerProvider = &TracerProvider{}

// NewTracerProvider returns a new and configured TracerProvider.
//
// By default the returned TracerProvider is configured with:
//  - a ParentBased(AlwaysSample) Sampler
//  - a random number IDGenerator
//  - the resource.Default() Resource
//  - the default SpanLimits.
//
// The passed opts are used to override these default values and configure the
// returned TracerProvider appropriately.
func NewTracerProvider(opts ...TracerProviderOption) *TracerProvider {
	o := &TracerProviderConfig{}

	for _, opt := range opts {
		opt(o)
	}

	ensureValidTracerProviderConfig(o)

	tp := &TracerProvider{
		namedTracer: make(map[instrumentation.Library]*tracer),
		sampler:     o.sampler,
		idGenerator: o.idGenerator,
		spanLimits:  o.spanLimits,
		resource:    o.resource,
	}
<<<<<<< HEAD

||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

	global.Info("TracerProvider created", "config", o)

=======
	global.Info("TracerProvider created", "config", o)

	spss := make(spanProcessorStates, 0, len(o.processors))
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	for _, sp := range o.processors {
		spss = append(spss, newSpanProcessorState(sp))
	}
	tp.spanProcessors.Store(&spss)

	return tp
}

// Tracer returns a Tracer with the given name and options. If a Tracer for
// the given name and options does not exist it is created, otherwise the
// existing Tracer is returned.
//
// If name is empty, DefaultTracerName is used instead.
//
// This method is safe to be called concurrently.
func (p *TracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	// This check happens before the mutex is acquired to avoid deadlocking if Tracer() is called from within Shutdown().
	if p.isShutdown.Load() {
		return noop.NewTracerProvider().Tracer(name, opts...)
	}
	c := trace.NewTracerConfig(opts...)
	if name == "" {
		name = defaultTracerName
	}
	il := instrumentation.Library{
		Name:    name,
		Version: c.InstrumentationVersion,
	}
<<<<<<< HEAD
	t, ok := p.namedTracer[il]
	if !ok {
		t = &tracer{
			provider:               p,
			instrumentationLibrary: il,
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	t, ok := p.namedTracer[is]
	if !ok {
		t = &tracer{
			provider:             p,
			instrumentationScope: is,
=======

	t, ok := func() (trace.Tracer, bool) {
		p.mu.Lock()
		defer p.mu.Unlock()
		// Must check the flag after acquiring the mutex to avoid returning a valid tracer if Shutdown() ran
		// after the first check above but before we acquired the mutex.
		if p.isShutdown.Load() {
			return noop.NewTracerProvider().Tracer(name, opts...), true
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		}
<<<<<<< HEAD
		p.namedTracer[il] = t
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		p.namedTracer[is] = t
		global.Info("Tracer created", "name", name, "version", c.InstrumentationVersion(), "schemaURL", c.SchemaURL())
=======
		t, ok := p.namedTracer[is]
		if !ok {
			t = &tracer{
				provider:             p,
				instrumentationScope: is,
			}
			p.namedTracer[is] = t
		}
		return t, ok
	}()
	if !ok {
		// This code is outside the mutex to not hold the lock while calling third party logging code:
		// - That code may do slow things like I/O, which would prolong the duration the lock is held,
		//   slowing down all tracing consumers.
		// - Logging code may be instrumented with tracing and deadlock because it could try
		//   acquiring the same non-reentrant mutex.
		global.Info("Tracer created", "name", name, "version", is.Version, "schemaURL", is.SchemaURL)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	}
	return t
}

<<<<<<< HEAD
// RegisterSpanProcessor adds the given SpanProcessor to the list of SpanProcessors
func (p *TracerProvider) RegisterSpanProcessor(s SpanProcessor) {
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// RegisterSpanProcessor adds the given SpanProcessor to the list of SpanProcessors.
func (p *TracerProvider) RegisterSpanProcessor(s SpanProcessor) {
=======
// RegisterSpanProcessor adds the given SpanProcessor to the list of SpanProcessors.
func (p *TracerProvider) RegisterSpanProcessor(sp SpanProcessor) {
	// This check prevents calls during a shutdown.
	if p.isShutdown.Load() {
		return
	}
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	p.mu.Lock()
	defer p.mu.Unlock()
<<<<<<< HEAD
	new := spanProcessorStates{}
	if old, ok := p.spanProcessors.Load().(spanProcessorStates); ok {
		new = append(new, old...)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	newSPS := spanProcessorStates{}
	if old, ok := p.spanProcessors.Load().(spanProcessorStates); ok {
		newSPS = append(newSPS, old...)
=======
	// This check prevents calls after a shutdown.
	if p.isShutdown.Load() {
		return
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	}
<<<<<<< HEAD
	newSpanSync := &spanProcessorState{
		sp:    s,
		state: &sync.Once{},
	}
	new = append(new, newSpanSync)
	p.spanProcessors.Store(new)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	newSpanSync := &spanProcessorState{
		sp:    s,
		state: &sync.Once{},
	}
	newSPS = append(newSPS, newSpanSync)
	p.spanProcessors.Store(newSPS)
=======

	current := p.getSpanProcessors()
	newSPS := make(spanProcessorStates, 0, len(current)+1)
	newSPS = append(newSPS, current...)
	newSPS = append(newSPS, newSpanProcessorState(sp))
	p.spanProcessors.Store(&newSPS)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

<<<<<<< HEAD
// UnregisterSpanProcessor removes the given SpanProcessor from the list of SpanProcessors
func (p *TracerProvider) UnregisterSpanProcessor(s SpanProcessor) {
	p.mu.Lock()
	defer p.mu.Unlock()
	spss := spanProcessorStates{}
	old, ok := p.spanProcessors.Load().(spanProcessorStates)
	if !ok || len(old) == 0 {
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// UnregisterSpanProcessor removes the given SpanProcessor from the list of SpanProcessors.
func (p *TracerProvider) UnregisterSpanProcessor(s SpanProcessor) {
	p.mu.Lock()
	defer p.mu.Unlock()
	spss := spanProcessorStates{}
	old, ok := p.spanProcessors.Load().(spanProcessorStates)
	if !ok || len(old) == 0 {
=======
// UnregisterSpanProcessor removes the given SpanProcessor from the list of SpanProcessors.
func (p *TracerProvider) UnregisterSpanProcessor(sp SpanProcessor) {
	// This check prevents calls during a shutdown.
	if p.isShutdown.Load() {
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	// This check prevents calls after a shutdown.
	if p.isShutdown.Load() {
		return
	}
	old := p.getSpanProcessors()
	if len(old) == 0 {
		return
	}
	spss := make(spanProcessorStates, len(old))
	copy(spss, old)

	// stop the span processor if it is started and remove it from the list
	var stopOnce *spanProcessorState
	var idx int
	for i, sps := range spss {
		if sps.sp == sp {
			stopOnce = sps
			idx = i
		}
	}
	if stopOnce != nil {
		stopOnce.state.Do(func() {
			if err := sp.Shutdown(context.Background()); err != nil {
				otel.Handle(err)
			}
		})
	}
	if len(spss) > 1 {
		copy(spss[idx:], spss[idx+1:])
	}
	spss[len(spss)-1] = nil
	spss = spss[:len(spss)-1]

	p.spanProcessors.Store(&spss)
}

// ForceFlush immediately exports all spans that have not yet been exported for
// all the registered span processors.
func (p *TracerProvider) ForceFlush(ctx context.Context) error {
	spss := p.getSpanProcessors()
	if len(spss) == 0 {
		return nil
	}

	for _, sps := range spss {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := sps.sp.ForceFlush(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Shutdown shuts down TracerProvider. All registered span processors are shut down
// in the order they were registered and any held computational resources are released.
// After Shutdown is called, all methods are no-ops.
func (p *TracerProvider) Shutdown(ctx context.Context) error {
	// This check prevents deadlocks in case of recursive shutdown.
	if p.isShutdown.Load() {
		return nil
	}
<<<<<<< HEAD
	if len(spss) == 0 {
		return nil
	}

	for _, sps := range spss {
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	var retErr error
	for _, sps := range spss {
=======
	p.mu.Lock()
	defer p.mu.Unlock()
	// This check prevents calls after a shutdown has already been done concurrently.
	if !p.isShutdown.CompareAndSwap(false, true) { // did toggle?
		return nil
	}

	var retErr error
	for _, sps := range p.getSpanProcessors() {
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		var err error
		sps.state.Do(func() {
			err = sps.sp.Shutdown(ctx)
		})
		if err != nil {
			return err
		}
	}
<<<<<<< HEAD
	return nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return retErr
}

// TracerProviderOption configures a TracerProvider.
type TracerProviderOption interface {
	apply(tracerProviderConfig) tracerProviderConfig
}

type traceProviderOptionFunc func(tracerProviderConfig) tracerProviderConfig

func (fn traceProviderOptionFunc) apply(cfg tracerProviderConfig) tracerProviderConfig {
	return fn(cfg)
=======
	p.spanProcessors.Store(&spanProcessorStates{})
	return retErr
}

func (p *TracerProvider) getSpanProcessors() spanProcessorStates {
	return *(p.spanProcessors.Load())
}

// TracerProviderOption configures a TracerProvider.
type TracerProviderOption interface {
	apply(tracerProviderConfig) tracerProviderConfig
}

type traceProviderOptionFunc func(tracerProviderConfig) tracerProviderConfig

func (fn traceProviderOptionFunc) apply(cfg tracerProviderConfig) tracerProviderConfig {
	return fn(cfg)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// WithSyncer registers the exporter with the TracerProvider using a
// SimpleSpanProcessor.
func WithSyncer(e SpanExporter) TracerProviderOption {
	return WithSpanProcessor(NewSimpleSpanProcessor(e))
}

// WithBatcher registers the exporter with the TracerProvider using a
// BatchSpanProcessor configured with the passed opts.
func WithBatcher(e SpanExporter, opts ...BatchSpanProcessorOption) TracerProviderOption {
	return WithSpanProcessor(NewBatchSpanProcessor(e, opts...))
}

// WithSpanProcessor registers the SpanProcessor with a TracerProvider.
func WithSpanProcessor(sp SpanProcessor) TracerProviderOption {
	return func(opts *TracerProviderConfig) {
		opts.processors = append(opts.processors, sp)
	}
}

// WithResource returns a TracerProviderOption that will configure the
// Resource r as a TracerProvider's Resource. The configured Resource is
// referenced by all the Tracers the TracerProvider creates. It represents the
// entity producing telemetry.
//
// If this option is not used, the TracerProvider will use the
// resource.Default() Resource by default.
func WithResource(r *resource.Resource) TracerProviderOption {
	return func(opts *TracerProviderConfig) {
		opts.resource = resource.Merge(resource.Environment(), r)
	}
}

// WithIDGenerator returns a TracerProviderOption that will configure the
// IDGenerator g as a TracerProvider's IDGenerator. The configured IDGenerator
// is used by the Tracers the TracerProvider creates to generate new Span and
// Trace IDs.
//
// If this option is not used, the TracerProvider will use a random number
// IDGenerator by default.
func WithIDGenerator(g IDGenerator) TracerProviderOption {
	return func(opts *TracerProviderConfig) {
		if g != nil {
			opts.idGenerator = g
		}
	}
}

// WithSampler returns a TracerProviderOption that will configure the Sampler
// s as a TracerProvider's Sampler. The configured Sampler is used by the
// Tracers the TracerProvider creates to make their sampling decisions for the
// Spans they create.
//
// If this option is not used, the TracerProvider will use a
// ParentBased(AlwaysSample) Sampler by default.
func WithSampler(s Sampler) TracerProviderOption {
	return func(opts *TracerProviderConfig) {
		if s != nil {
			opts.sampler = s
		}
	}
}

// WithSpanLimits returns a TracerProviderOption that will configure the
// SpanLimits sl as a TracerProvider's SpanLimits. The configured SpanLimits
// are used used by the Tracers the TracerProvider and the Spans they create
// to limit tracing resources used.
//
// If this option is not used, the TracerProvider will use the default
// SpanLimits.
func WithSpanLimits(sl SpanLimits) TracerProviderOption {
	return func(opts *TracerProviderConfig) {
		opts.spanLimits = sl
	}
}

// ensureValidTracerProviderConfig ensures that given TracerProviderConfig is valid.
func ensureValidTracerProviderConfig(cfg *TracerProviderConfig) {
	if cfg.sampler == nil {
		cfg.sampler = ParentBased(AlwaysSample())
	}
	if cfg.idGenerator == nil {
		cfg.idGenerator = defaultIDGenerator()
	}
	cfg.spanLimits.ensureDefault()
	if cfg.resource == nil {
		cfg.resource = resource.Default()
	}
}
