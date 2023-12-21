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

package metric // import "go.opentelemetry.io/otel/sdk/metric"

<<<<<<< HEAD:vendor/go.opentelemetry.io/otel/sdk/metric/atomicfields.go
import "unsafe"
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/version.go
// Version is the current release version of the gRPC instrumentation.
func Version() string {
	return "0.35.0"
	// This string is updated by the pre_release.sh script during release
}
=======
// Version is the current release version of the gRPC instrumentation.
func Version() string {
	return "0.46.1"
	// This string is updated by the pre_release.sh script during release
}
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/version.go

<<<<<<< HEAD:vendor/go.opentelemetry.io/otel/sdk/metric/atomicfields.go
func AtomicFieldOffsets() map[string]uintptr {
	return map[string]uintptr{
		"record.refMapped.value": unsafe.Offsetof(record{}.refMapped.value),
		"record.updateCount":     unsafe.Offsetof(record{}.updateCount),
	}
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/version.go
// SemVersion is the semantic version to be supplied to tracer/meter creation.
func SemVersion() string {
	return "semver:" + Version()
=======
// SemVersion is the semantic version to be supplied to tracer/meter creation.
//
// Deprecated: Use [Version] instead.
func SemVersion() string {
	return Version()
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/version.go
}
