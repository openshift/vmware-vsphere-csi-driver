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

package resource // import "go.opentelemetry.io/otel/sdk/resource"

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
=======
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
)

<<<<<<< HEAD
// envVar is the environment variable name OpenTelemetry Resource information can be assigned to.
const envVar = "OTEL_RESOURCE_ATTRIBUTES"
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
const (
	// resourceAttrKey is the environment variable name OpenTelemetry Resource information will be read from.
	resourceAttrKey = "OTEL_RESOURCE_ATTRIBUTES"

	// svcNameKey is the environment variable name that Service Name information will be read from.
	svcNameKey = "OTEL_SERVICE_NAME"
)
=======
const (
	// resourceAttrKey is the environment variable name OpenTelemetry Resource information will be read from.
	resourceAttrKey = "OTEL_RESOURCE_ATTRIBUTES" //nolint:gosec // False positive G101: Potential hardcoded credentials

	// svcNameKey is the environment variable name that Service Name information will be read from.
	svcNameKey = "OTEL_SERVICE_NAME"
)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

// errMissingValue is returned when a resource value is missing.
var errMissingValue = fmt.Errorf("%w: missing value", ErrPartialResource)

// FromEnv is a Detector that implements the Detector and collects
// resources from environment.  This Detector is included as a
// builtin.  If these resource attributes are not wanted, use the
// WithFromEnv(nil) or WithoutBuiltin() options to explicitly disable
// them.
type FromEnv struct{}

// compile time assertion that FromEnv implements Detector interface
var _ Detector = FromEnv{}

// Detect collects resources from environment
func (FromEnv) Detect(context.Context) (*Resource, error) {
	attrs := strings.TrimSpace(os.Getenv(envVar))

	if attrs == "" {
		return Empty(), nil
	}
<<<<<<< HEAD
	return constructOTResources(attrs)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

	var res *Resource

	if svcName != "" {
		res = NewSchemaless(semconv.ServiceNameKey.String(svcName))
	}

	r2, err := constructOTResources(attrs)

	// Ensure that the resource with the service name from OTEL_SERVICE_NAME
	// takes precedence, if it was defined.
	res, err2 := Merge(r2, res)

	if err == nil {
		err = err2
	} else if err2 != nil {
		err = fmt.Errorf("detecting resources: %s", []string{err.Error(), err2.Error()})
	}

	return res, err
=======

	var res *Resource

	if svcName != "" {
		res = NewSchemaless(semconv.ServiceName(svcName))
	}

	r2, err := constructOTResources(attrs)

	// Ensure that the resource with the service name from OTEL_SERVICE_NAME
	// takes precedence, if it was defined.
	res, err2 := Merge(r2, res)

	if err == nil {
		err = err2
	} else if err2 != nil {
		err = fmt.Errorf("detecting resources: %s", []string{err.Error(), err2.Error()})
	}

	return res, err
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

func constructOTResources(s string) (*Resource, error) {
	pairs := strings.Split(s, ",")
	var attrs []attribute.KeyValue
	var invalid []string
	for _, p := range pairs {
		k, v, found := strings.Cut(p, "=")
		if !found {
			invalid = append(invalid, p)
			continue
		}
		key := strings.TrimSpace(k)
		val, err := url.PathUnescape(strings.TrimSpace(v))
		if err != nil {
			// Retain original value if decoding fails, otherwise it will be
			// an empty string.
			val = v
			otel.Handle(err)
		}
		attrs = append(attrs, attribute.String(key, val))
	}
	var err error
	if len(invalid) > 0 {
		err = fmt.Errorf("%w: %v", errMissingValue, invalid)
	}
	return NewWithAttributes(attrs...), err
}
