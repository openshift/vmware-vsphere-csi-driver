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
	"strings"

<<<<<<< HEAD
	"go.opentelemetry.io/otel/semconv"
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
=======
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
)

<<<<<<< HEAD
type osTypeDetector struct{}
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
type osDescriptionProvider func() (string, error)

var defaultOSDescriptionProvider osDescriptionProvider = platformOSDescription

var osDescription = defaultOSDescriptionProvider

func setDefaultOSDescriptionProvider() {
	setOSDescriptionProvider(defaultOSDescriptionProvider)
}

func setOSDescriptionProvider(osDescriptionProvider osDescriptionProvider) {
	osDescription = osDescriptionProvider
}

type osTypeDetector struct{}
type osDescriptionDetector struct{}
=======
type osDescriptionProvider func() (string, error)

var defaultOSDescriptionProvider osDescriptionProvider = platformOSDescription

var osDescription = defaultOSDescriptionProvider

func setDefaultOSDescriptionProvider() {
	setOSDescriptionProvider(defaultOSDescriptionProvider)
}

func setOSDescriptionProvider(osDescriptionProvider osDescriptionProvider) {
	osDescription = osDescriptionProvider
}

type (
	osTypeDetector        struct{}
	osDescriptionDetector struct{}
)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

// Detect returns a *Resource that describes the operating system type the
// service is running on.
func (osTypeDetector) Detect(ctx context.Context) (*Resource, error) {
	osType := runtimeOS()

	return NewWithAttributes(
		semconv.OSTypeKey.String(strings.ToLower(osType)),
	), nil
}

<<<<<<< HEAD
// WithOSType adds an attribute with the operating system type to the configured Resource.
func WithOSType() Option {
	return WithDetectors(osTypeDetector{})
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
// Detect returns a *Resource that describes the operating system the
// service is running on.
func (osDescriptionDetector) Detect(ctx context.Context) (*Resource, error) {
	description, err := osDescription()

	if err != nil {
		return nil, err
	}

	return NewWithAttributes(
		semconv.SchemaURL,
		semconv.OSDescriptionKey.String(description),
	), nil
}

// mapRuntimeOSToSemconvOSType translates the OS name as provided by the Go runtime
// into an OS type attribute with the corresponding value defined by the semantic
// conventions. In case the provided OS name isn't mapped, it's transformed to lowercase
// and used as the value for the returned OS type attribute.
func mapRuntimeOSToSemconvOSType(osType string) attribute.KeyValue {
	// the elements in this map are the intersection between
	// available GOOS values and defined semconv OS types
	osTypeAttributeMap := map[string]attribute.KeyValue{
		"darwin":    semconv.OSTypeDarwin,
		"dragonfly": semconv.OSTypeDragonflyBSD,
		"freebsd":   semconv.OSTypeFreeBSD,
		"linux":     semconv.OSTypeLinux,
		"netbsd":    semconv.OSTypeNetBSD,
		"openbsd":   semconv.OSTypeOpenBSD,
		"solaris":   semconv.OSTypeSolaris,
		"windows":   semconv.OSTypeWindows,
	}

	var osTypeAttribute attribute.KeyValue

	if attr, ok := osTypeAttributeMap[osType]; ok {
		osTypeAttribute = attr
	} else {
		osTypeAttribute = semconv.OSTypeKey.String(strings.ToLower(osType))
	}

	return osTypeAttribute
=======
// Detect returns a *Resource that describes the operating system the
// service is running on.
func (osDescriptionDetector) Detect(ctx context.Context) (*Resource, error) {
	description, err := osDescription()
	if err != nil {
		return nil, err
	}

	return NewWithAttributes(
		semconv.SchemaURL,
		semconv.OSDescription(description),
	), nil
}

// mapRuntimeOSToSemconvOSType translates the OS name as provided by the Go runtime
// into an OS type attribute with the corresponding value defined by the semantic
// conventions. In case the provided OS name isn't mapped, it's transformed to lowercase
// and used as the value for the returned OS type attribute.
func mapRuntimeOSToSemconvOSType(osType string) attribute.KeyValue {
	// the elements in this map are the intersection between
	// available GOOS values and defined semconv OS types
	osTypeAttributeMap := map[string]attribute.KeyValue{
		"aix":       semconv.OSTypeAIX,
		"darwin":    semconv.OSTypeDarwin,
		"dragonfly": semconv.OSTypeDragonflyBSD,
		"freebsd":   semconv.OSTypeFreeBSD,
		"linux":     semconv.OSTypeLinux,
		"netbsd":    semconv.OSTypeNetBSD,
		"openbsd":   semconv.OSTypeOpenBSD,
		"solaris":   semconv.OSTypeSolaris,
		"windows":   semconv.OSTypeWindows,
		"zos":       semconv.OSTypeZOS,
	}

	var osTypeAttribute attribute.KeyValue

	if attr, ok := osTypeAttributeMap[osType]; ok {
		osTypeAttribute = attr
	} else {
		osTypeAttribute = semconv.OSTypeKey.String(strings.ToLower(osType))
	}

	return osTypeAttribute
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}
