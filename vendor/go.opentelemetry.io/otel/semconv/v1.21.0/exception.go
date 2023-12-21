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

<<<<<<<< HEAD:vendor/go.opentelemetry.io/otel/unit/unit.go
package unit // import "go.opentelemetry.io/otel/unit"
|||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/otel/metric/unit/unit.go
package unit // import "go.opentelemetry.io/otel/metric/unit"
========
package semconv // import "go.opentelemetry.io/otel/semconv/v1.21.0"
>>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/otel/semconv/v1.21.0/exception.go

<<<<<<<< HEAD:vendor/go.opentelemetry.io/otel/unit/unit.go
type Unit string

|||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/otel/metric/unit/unit.go
// Unit is a determinate standard quantity of measurement.
type Unit string

// Units defined by OpenTelemetry.
========
>>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686)):vendor/go.opentelemetry.io/otel/semconv/v1.21.0/exception.go
const (
	// ExceptionEventName is the name of the Span event representing an exception.
	ExceptionEventName = "exception"
)
