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

package otelgrpc

// gRPC tracing middleware
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/rpc.md
import (
	"context"
	"io"
	"net"
<<<<<<< HEAD
	"strings"

	"github.com/golang/protobuf/proto" // nolint:staticcheck
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

	"github.com/golang/protobuf/proto" // nolint:staticcheck
=======
	"strconv"
	"time"
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

	"google.golang.org/grpc"
	grpc_codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
<<<<<<< HEAD
	"go.opentelemetry.io/otel/semconv"
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
=======
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	"go.opentelemetry.io/otel/trace"

	otelcontrib "go.opentelemetry.io/contrib"
)

type messageType attribute.KeyValue

// Event adds an event of the messageType to the span associated with the
// passed context with a message id.
func (m messageType) Event(ctx context.Context, id int, _ interface{}) {
	span := trace.SpanFromContext(ctx)
<<<<<<< HEAD
	if p, ok := message.(proto.Message); ok {
		span.AddEvent("message", trace.WithAttributes(
			attribute.KeyValue(m),
			semconv.RPCMessageIDKey.Int(id),
			semconv.RPCMessageUncompressedSizeKey.Int(proto.Size(p)),
		))
	} else {
		span.AddEvent("message", trace.WithAttributes(
			attribute.KeyValue(m),
			semconv.RPCMessageIDKey.Int(id),
		))
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	if p, ok := message.(proto.Message); ok {
		span.AddEvent("message", trace.WithAttributes(
			attribute.KeyValue(m),
			RPCMessageIDKey.Int(id),
			RPCMessageUncompressedSizeKey.Int(proto.Size(p)),
		))
	} else {
		span.AddEvent("message", trace.WithAttributes(
			attribute.KeyValue(m),
			RPCMessageIDKey.Int(id),
		))
=======
	if !span.IsRecording() {
		return
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	}
	span.AddEvent("message", trace.WithAttributes(
		attribute.KeyValue(m),
		RPCMessageIDKey.Int(id),
	))
}

var (
	messageSent     = messageType(semconv.RPCMessageTypeSent)
	messageReceived = messageType(semconv.RPCMessageTypeReceived)
)

// UnaryClientInterceptor returns a grpc.UnaryClientInterceptor suitable
// for use in a grpc.Dial call.
//
// Deprecated: Use [NewClientHandler] instead.
func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	cfg := newConfig(opts)
=======
	cfg := newConfig(opts, "client")
	tracer := cfg.TracerProvider.Tracer(
		ScopeName,
		trace.WithInstrumentationVersion(Version()),
	)

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		callOpts ...grpc.CallOption,
	) error {
<<<<<<< HEAD
		requestMetadata, _ := metadata.FromOutgoingContext(ctx)
		metadataCopy := requestMetadata.Copy()
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		i := &InterceptorInfo{
			Method: method,
			Type:   UnaryClient,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return invoker(ctx, method, req, reply, cc, callOpts...)
		}

		requestMetadata, _ := metadata.FromOutgoingContext(ctx)
		metadataCopy := requestMetadata.Copy()
=======
		i := &InterceptorInfo{
			Method: method,
			Type:   UnaryClient,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return invoker(ctx, method, req, reply, cc, callOpts...)
		}

		name, attr, _ := telemetryAttributes(method, cc.Target())
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

<<<<<<< HEAD
		tracer := newConfig(opts).TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(otelcontrib.SemVersion()),
		)

		name, attr := spanInfo(method, cc.Target())
		var span trace.Span
		ctx, span = tracer.Start(
			ctx,
			name,
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		tracer := cfg.TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(SemVersion()),
		)

		name, attr := spanInfo(method, cc.Target())
		var span trace.Span
		ctx, span = tracer.Start(
			ctx,
			name,
=======
		startOpts := append([]trace.SpanStartOption{
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(attr...),
		},
			cfg.SpanStartOptions...,
		)

		ctx, span := tracer.Start(
			ctx,
			name,
			startOpts...,
		)
		defer span.End()

<<<<<<< HEAD
		Inject(ctx, &metadataCopy, opts...)
		ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		inject(ctx, &metadataCopy, cfg.Propagators)
		ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
=======
		ctx = inject(ctx, cfg.Propagators)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

		if cfg.SentEvent {
			messageSent.Event(ctx, 1, req)
		}

		err := invoker(ctx, method, req, reply, cc, callOpts...)

		if cfg.ReceivedEvent {
			messageReceived.Event(ctx, 1, reply)
		}

		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(statusCodeAttr(s.Code()))
		} else {
			span.SetAttributes(statusCodeAttr(grpc_codes.OK))
		}

		return err
	}
}

<<<<<<< HEAD
type streamEventType int

type streamEvent struct {
	Type streamEventType
	Err  error
}

const (
	closeEvent streamEventType = iota
	receiveEndEvent
	errorEvent
)

||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
type streamEventType int

type streamEvent struct {
	Type streamEventType
	Err  error
}

const (
	receiveEndEvent streamEventType = iota
	errorEvent
)

=======
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// clientStream  wraps around the embedded grpc.ClientStream, and intercepts the RecvMsg and
// SendMsg method call.
type clientStream struct {
	grpc.ClientStream
	desc *grpc.StreamDesc

	span trace.Span

	receivedEvent bool
	sentEvent     bool

	receivedMessageID int
	sentMessageID     int
}

var _ = proto.Marshal

func (w *clientStream) RecvMsg(m interface{}) error {
	err := w.ClientStream.RecvMsg(m)

	if err == nil && !w.desc.ServerStreams {
		w.endSpan(nil)
	} else if err == io.EOF {
		w.endSpan(nil)
	} else if err != nil {
		w.endSpan(err)
	} else {
		w.receivedMessageID++

		if w.receivedEvent {
			messageReceived.Event(w.Context(), w.receivedMessageID, m)
		}
	}

	return err
}

func (w *clientStream) SendMsg(m interface{}) error {
	err := w.ClientStream.SendMsg(m)

	w.sentMessageID++

	if w.sentEvent {
		messageSent.Event(w.Context(), w.sentMessageID, m)
	}

	if err != nil {
		w.endSpan(err)
	}

	return err
}

func (w *clientStream) Header() (metadata.MD, error) {
	md, err := w.ClientStream.Header()
	if err != nil {
		w.endSpan(err)
	}

	return md, err
}

func (w *clientStream) CloseSend() error {
	err := w.ClientStream.CloseSend()
	if err != nil {
<<<<<<< HEAD
		w.sendStreamEvent(errorEvent, err)
	} else {
		w.sendStreamEvent(closeEvent, nil)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		w.sendStreamEvent(errorEvent, err)
=======
		w.endSpan(err)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	}

	return err
}

<<<<<<< HEAD
const (
	clientClosedState byte = 1 << iota
	receiveEndedState
)

func wrapClientStream(s grpc.ClientStream, desc *grpc.StreamDesc) *clientStream {
	events := make(chan streamEvent)
	eventsDone := make(chan struct{})
	finished := make(chan error)

	go func() {
		defer close(eventsDone)

		// Both streams have to be closed
		state := byte(0)

		for event := range events {
			switch event.Type {
			case closeEvent:
				state |= clientClosedState
			case receiveEndEvent:
				state |= receiveEndedState
			case errorEvent:
				finished <- event.Err
				return
			}

			if state == clientClosedState|receiveEndedState {
				finished <- nil
				return
			}
		}
	}()

||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
func wrapClientStream(ctx context.Context, s grpc.ClientStream, desc *grpc.StreamDesc) *clientStream {
	events := make(chan streamEvent)
	eventsDone := make(chan struct{})
	finished := make(chan error)

	go func() {
		defer close(eventsDone)

		for {
			select {
			case event := <-events:
				switch event.Type {
				case receiveEndEvent:
					finished <- nil
					return
				case errorEvent:
					finished <- event.Err
					return
				}
			case <-ctx.Done():
				finished <- ctx.Err()
				return
			}
		}
	}()

=======
func wrapClientStream(ctx context.Context, s grpc.ClientStream, desc *grpc.StreamDesc, span trace.Span, cfg *config) *clientStream {
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return &clientStream{
		ClientStream:  s,
		span:          span,
		desc:          desc,
		receivedEvent: cfg.ReceivedEvent,
		sentEvent:     cfg.SentEvent,
	}
}

func (w *clientStream) endSpan(err error) {
	if err != nil {
		s, _ := status.FromError(err)
		w.span.SetStatus(codes.Error, s.Message())
		w.span.SetAttributes(statusCodeAttr(s.Code()))
	} else {
		w.span.SetAttributes(statusCodeAttr(grpc_codes.OK))
	}

	w.span.End()
}

// StreamClientInterceptor returns a grpc.StreamClientInterceptor suitable
// for use in a grpc.Dial call.
//
// Deprecated: Use [NewClientHandler] instead.
func StreamClientInterceptor(opts ...Option) grpc.StreamClientInterceptor {
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	cfg := newConfig(opts)
=======
	cfg := newConfig(opts, "client")
	tracer := cfg.TracerProvider.Tracer(
		ScopeName,
		trace.WithInstrumentationVersion(Version()),
	)

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		callOpts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
<<<<<<< HEAD
		requestMetadata, _ := metadata.FromOutgoingContext(ctx)
		metadataCopy := requestMetadata.Copy()
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		i := &InterceptorInfo{
			Method: method,
			Type:   StreamClient,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return streamer(ctx, desc, cc, method, callOpts...)
		}

		requestMetadata, _ := metadata.FromOutgoingContext(ctx)
		metadataCopy := requestMetadata.Copy()
=======
		i := &InterceptorInfo{
			Method: method,
			Type:   StreamClient,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return streamer(ctx, desc, cc, method, callOpts...)
		}

		name, attr, _ := telemetryAttributes(method, cc.Target())
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

<<<<<<< HEAD
		tracer := newConfig(opts).TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(otelcontrib.SemVersion()),
		)

		name, attr := spanInfo(method, cc.Target())
		var span trace.Span
		ctx, span = tracer.Start(
			ctx,
			name,
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		tracer := cfg.TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(SemVersion()),
		)

		name, attr := spanInfo(method, cc.Target())
		var span trace.Span
		ctx, span = tracer.Start(
			ctx,
			name,
=======
		startOpts := append([]trace.SpanStartOption{
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(attr...),
		},
			cfg.SpanStartOptions...,
		)

<<<<<<< HEAD
		Inject(ctx, &metadataCopy, opts...)
		ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		inject(ctx, &metadataCopy, cfg.Propagators)
		ctx = metadata.NewOutgoingContext(ctx, metadataCopy)
=======
		ctx, span := tracer.Start(
			ctx,
			name,
			startOpts...,
		)

		ctx = inject(ctx, cfg.Propagators)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

		s, err := streamer(ctx, desc, cc, method, callOpts...)
		if err != nil {
			grpcStatus, _ := status.FromError(err)
			span.SetStatus(codes.Error, grpcStatus.Message())
			span.SetAttributes(statusCodeAttr(grpcStatus.Code()))
			span.End()
			return s, err
		}
<<<<<<< HEAD
		stream := wrapClientStream(s, desc)

		go func() {
			err := <-stream.finished

			if err != nil {
				s, _ := status.FromError(err)
				span.SetStatus(codes.Error, s.Message())
				span.SetAttributes(statusCodeAttr(s.Code()))
			} else {
				span.SetAttributes(statusCodeAttr(grpc_codes.OK))
			}

			span.End()
		}()

||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		stream := wrapClientStream(ctx, s, desc)

		go func() {
			err := <-stream.finished

			if err != nil {
				s, _ := status.FromError(err)
				span.SetStatus(codes.Error, s.Message())
				span.SetAttributes(statusCodeAttr(s.Code()))
			} else {
				span.SetAttributes(statusCodeAttr(grpc_codes.OK))
			}

			span.End()
		}()

=======
		stream := wrapClientStream(ctx, s, desc, span, cfg)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		return stream, nil
	}
}

// UnaryServerInterceptor returns a grpc.UnaryServerInterceptor suitable
// for use in a grpc.NewServer call.
//
// Deprecated: Use [NewServerHandler] instead.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	cfg := newConfig(opts)
=======
	cfg := newConfig(opts, "server")
	tracer := cfg.TracerProvider.Tracer(
		ScopeName,
		trace.WithInstrumentationVersion(Version()),
	)

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
<<<<<<< HEAD
		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		i := &InterceptorInfo{
			UnaryServerInfo: info,
			Type:            UnaryServer,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return handler(ctx, req)
		}

		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()
=======
		i := &InterceptorInfo{
			UnaryServerInfo: info,
			Type:            UnaryServer,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return handler(ctx, req)
		}

		ctx = extract(ctx, cfg.Propagators)
		name, attr, metricAttrs := telemetryAttributes(info.FullMethod, peerFromCtx(ctx))
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

<<<<<<< HEAD
		entries, spanCtx := Extract(ctx, &metadataCopy, opts...)
		ctx = baggage.ContextWithValues(ctx, entries...)

		tracer := newConfig(opts).TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(otelcontrib.SemVersion()),
		)

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx))
		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		bags, spanCtx := Extract(ctx, &metadataCopy, opts...)
		ctx = baggage.ContextWithBaggage(ctx, bags)

		tracer := cfg.TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(SemVersion()),
		)

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx))
		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
=======
		startOpts := append([]trace.SpanStartOption{
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attr...),
		},
			cfg.SpanStartOptions...,
		)

		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, trace.SpanContextFromContext(ctx)),
			name,
			startOpts...,
		)
		defer span.End()

		if cfg.ReceivedEvent {
			messageReceived.Event(ctx, 1, req)
		}

		before := time.Now()

		resp, err := handler(ctx, req)

		s, _ := status.FromError(err)
		if err != nil {
			statusCode, msg := serverStatus(s)
			span.SetStatus(statusCode, msg)
			if cfg.SentEvent {
				messageSent.Event(ctx, 1, s.Proto())
			}
		} else {
			if cfg.SentEvent {
				messageSent.Event(ctx, 1, resp)
			}
		}
		grpcStatusCodeAttr := statusCodeAttr(s.Code())
		span.SetAttributes(grpcStatusCodeAttr)

		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedTime := float64(time.Since(before)) / float64(time.Millisecond)

		metricAttrs = append(metricAttrs, grpcStatusCodeAttr)
		cfg.rpcDuration.Record(ctx, elapsedTime, metric.WithAttributes(metricAttrs...))

		return resp, err
	}
}

// serverStream wraps around the embedded grpc.ServerStream, and intercepts the RecvMsg and
// SendMsg method call.
type serverStream struct {
	grpc.ServerStream
	ctx context.Context

	receivedMessageID int
	sentMessageID     int

	receivedEvent bool
	sentEvent     bool
}

func (w *serverStream) Context() context.Context {
	return w.ctx
}

func (w *serverStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)

	if err == nil {
		w.receivedMessageID++
		if w.receivedEvent {
			messageReceived.Event(w.Context(), w.receivedMessageID, m)
		}
	}

	return err
}

func (w *serverStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)

	w.sentMessageID++
	if w.sentEvent {
		messageSent.Event(w.Context(), w.sentMessageID, m)
	}

	return err
}

func wrapServerStream(ctx context.Context, ss grpc.ServerStream, cfg *config) *serverStream {
	return &serverStream{
		ServerStream:  ss,
		ctx:           ctx,
		receivedEvent: cfg.ReceivedEvent,
		sentEvent:     cfg.SentEvent,
	}
}

// StreamServerInterceptor returns a grpc.StreamServerInterceptor suitable
// for use in a grpc.NewServer call.
//
// Deprecated: Use [NewServerHandler] instead.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	cfg := newConfig(opts)
=======
	cfg := newConfig(opts, "server")
	tracer := cfg.TracerProvider.Tracer(
		ScopeName,
		trace.WithInstrumentationVersion(Version()),
	)

>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := ss.Context()
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		i := &InterceptorInfo{
			StreamServerInfo: info,
			Type:             StreamServer,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return handler(srv, wrapServerStream(ctx, ss))
		}
=======
		i := &InterceptorInfo{
			StreamServerInfo: info,
			Type:             StreamServer,
		}
		if cfg.Filter != nil && !cfg.Filter(i) {
			return handler(srv, wrapServerStream(ctx, ss, cfg))
		}
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

		ctx = extract(ctx, cfg.Propagators)
		name, attr, _ := telemetryAttributes(info.FullMethod, peerFromCtx(ctx))

<<<<<<< HEAD
		entries, spanCtx := Extract(ctx, &metadataCopy, opts...)
		ctx = baggage.ContextWithValues(ctx, entries...)

		tracer := newConfig(opts).TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(otelcontrib.SemVersion()),
		)

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx))
		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		bags, spanCtx := Extract(ctx, &metadataCopy, opts...)
		ctx = baggage.ContextWithBaggage(ctx, bags)

		tracer := cfg.TracerProvider.Tracer(
			instrumentationName,
			trace.WithInstrumentationVersion(SemVersion()),
		)

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx))
		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
=======
		startOpts := append([]trace.SpanStartOption{
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attr...),
		},
			cfg.SpanStartOptions...,
		)

		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, trace.SpanContextFromContext(ctx)),
			name,
			startOpts...,
		)
		defer span.End()

		err := handler(srv, wrapServerStream(ctx, ss, cfg))
		if err != nil {
			s, _ := status.FromError(err)
			statusCode, msg := serverStatus(s)
			span.SetStatus(statusCode, msg)
			span.SetAttributes(statusCodeAttr(s.Code()))
		} else {
			span.SetAttributes(statusCodeAttr(grpc_codes.OK))
		}

		return err
	}
}

<<<<<<< HEAD
// spanInfo returns a span name and all appropriate attributes from the gRPC
// method and peer address.
func spanInfo(fullMethod, peerAddress string) (string, []attribute.KeyValue) {
	attrs := []attribute.KeyValue{semconv.RPCSystemGRPC}
	name, mAttrs := parseFullMethod(fullMethod)
	attrs = append(attrs, mAttrs...)
	attrs = append(attrs, peerAttr(peerAddress)...)
	return name, attrs
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// spanInfo returns a span name and all appropriate attributes from the gRPC
// method and peer address.
func spanInfo(fullMethod, peerAddress string) (string, []attribute.KeyValue) {
	attrs := []attribute.KeyValue{RPCSystemGRPC}
	name, mAttrs := internal.ParseFullMethod(fullMethod)
	attrs = append(attrs, mAttrs...)
	attrs = append(attrs, peerAttr(peerAddress)...)
	return name, attrs
=======
// telemetryAttributes returns a span name and span and metric attributes from
// the gRPC method and peer address.
func telemetryAttributes(fullMethod, peerAddress string) (string, []attribute.KeyValue, []attribute.KeyValue) {
	name, methodAttrs := internal.ParseFullMethod(fullMethod)
	peerAttrs := peerAttr(peerAddress)

	attrs := make([]attribute.KeyValue, 0, 1+len(methodAttrs)+len(peerAttrs))
	attrs = append(attrs, RPCSystemGRPC)
	attrs = append(attrs, methodAttrs...)
	metricAttrs := attrs[:1+len(methodAttrs)]
	attrs = append(attrs, peerAttrs...)
	return name, attrs, metricAttrs
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// peerAttr returns attributes about the peer address.
func peerAttr(addr string) []attribute.KeyValue {
	host, p, err := net.SplitHostPort(addr)
	if err != nil {
		return nil
	}

	if host == "" {
		host = "127.0.0.1"
	}
	port, err := strconv.Atoi(p)
	if err != nil {
		return nil
	}

	var attr []attribute.KeyValue
	if ip := net.ParseIP(host); ip != nil {
		attr = []attribute.KeyValue{
			semconv.NetSockPeerAddr(host),
			semconv.NetSockPeerPort(port),
		}
	} else {
		attr = []attribute.KeyValue{
			semconv.NetPeerName(host),
			semconv.NetPeerPort(port),
		}
	}

	return attr
}

// peerFromCtx returns a peer address from a context, if one exists.
func peerFromCtx(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	return p.Addr.String()
}

// parseFullMethod returns a span name following the OpenTelemetry semantic
// conventions as well as all applicable span attribute.KeyValue attributes based
// on a gRPC's FullMethod.
func parseFullMethod(fullMethod string) (string, []attribute.KeyValue) {
	name := strings.TrimLeft(fullMethod, "/")
	parts := strings.SplitN(name, "/", 2)
	if len(parts) != 2 {
		// Invalid format, does not follow `/package.service/method`.
		return name, []attribute.KeyValue(nil)
	}

	var attrs []attribute.KeyValue
	if service := parts[0]; service != "" {
		attrs = append(attrs, semconv.RPCServiceKey.String(service))
	}
	if method := parts[1]; method != "" {
		attrs = append(attrs, semconv.RPCMethodKey.String(method))
	}
	return name, attrs
}

// statusCodeAttr returns status code attribute based on given gRPC code
func statusCodeAttr(c grpc_codes.Code) attribute.KeyValue {
	return GRPCStatusCodeKey.Int64(int64(c))
}

// serverStatus returns a span status code and message for a given gRPC
// status code. It maps specific gRPC status codes to a corresponding span
// status code and message. This function is intended for use on the server
// side of a gRPC connection.
//
// If the gRPC status code is Unknown, DeadlineExceeded, Unimplemented,
// Internal, Unavailable, or DataLoss, it returns a span status code of Error
// and the message from the gRPC status. Otherwise, it returns a span status
// code of Unset and an empty message.
func serverStatus(grpcStatus *status.Status) (codes.Code, string) {
	switch grpcStatus.Code() {
	case grpc_codes.Unknown,
		grpc_codes.DeadlineExceeded,
		grpc_codes.Unimplemented,
		grpc_codes.Internal,
		grpc_codes.Unavailable,
		grpc_codes.DataLoss:
		return codes.Error, grpcStatus.Message()
	default:
		return codes.Unset, ""
	}
}
