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
	"os"
	"os/user"
	"path/filepath"
	"runtime"

<<<<<<< HEAD
	"go.opentelemetry.io/otel/semconv"
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
=======
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
)

type (
	pidProvider            func() int
	executablePathProvider func() (string, error)
	commandArgsProvider    func() []string
	ownerProvider          func() (*user.User, error)
	runtimeNameProvider    func() string
	runtimeVersionProvider func() string
	runtimeOSProvider      func() string
	runtimeArchProvider    func() string
)

var (
	defaultPidProvider            pidProvider            = os.Getpid
	defaultExecutablePathProvider executablePathProvider = os.Executable
	defaultCommandArgsProvider    commandArgsProvider    = func() []string { return os.Args }
	defaultOwnerProvider          ownerProvider          = user.Current
	defaultRuntimeNameProvider    runtimeNameProvider    = func() string { return runtime.Compiler }
	defaultRuntimeVersionProvider runtimeVersionProvider = runtime.Version
	defaultRuntimeOSProvider      runtimeOSProvider      = func() string { return runtime.GOOS }
	defaultRuntimeArchProvider    runtimeArchProvider    = func() string { return runtime.GOARCH }
)

var (
	pid            = defaultPidProvider
	executablePath = defaultExecutablePathProvider
	commandArgs    = defaultCommandArgsProvider
	owner          = defaultOwnerProvider
	runtimeName    = defaultRuntimeNameProvider
	runtimeVersion = defaultRuntimeVersionProvider
	runtimeOS      = defaultRuntimeOSProvider
	runtimeArch    = defaultRuntimeArchProvider
)

func setDefaultOSProviders() {
	setOSProviders(
		defaultPidProvider,
		defaultExecutablePathProvider,
		defaultCommandArgsProvider,
	)
}

func setOSProviders(
	pidProvider pidProvider,
	executablePathProvider executablePathProvider,
	commandArgsProvider commandArgsProvider,
) {
	pid = pidProvider
	executablePath = executablePathProvider
	commandArgs = commandArgsProvider
}

func setDefaultRuntimeProviders() {
	setRuntimeProviders(
		defaultRuntimeNameProvider,
		defaultRuntimeVersionProvider,
		defaultRuntimeOSProvider,
		defaultRuntimeArchProvider,
	)
}

func setRuntimeProviders(
	runtimeNameProvider runtimeNameProvider,
	runtimeVersionProvider runtimeVersionProvider,
	runtimeOSProvider runtimeOSProvider,
	runtimeArchProvider runtimeArchProvider,
) {
	runtimeName = runtimeNameProvider
	runtimeVersion = runtimeVersionProvider
	runtimeOS = runtimeOSProvider
	runtimeArch = runtimeArchProvider
}

func setDefaultUserProviders() {
	setUserProviders(defaultOwnerProvider)
}

func setUserProviders(ownerProvider ownerProvider) {
	owner = ownerProvider
}

type (
	processPIDDetector                struct{}
	processExecutableNameDetector     struct{}
	processExecutablePathDetector     struct{}
	processCommandArgsDetector        struct{}
	processOwnerDetector              struct{}
	processRuntimeNameDetector        struct{}
	processRuntimeVersionDetector     struct{}
	processRuntimeDescriptionDetector struct{}
)

// Detect returns a *Resource that describes the process identifier (PID) of the
// executing process.
func (processPIDDetector) Detect(ctx context.Context) (*Resource, error) {
<<<<<<< HEAD
	return NewWithAttributes(semconv.ProcessPIDKey.Int(pid())), nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessPIDKey.Int(pid())), nil
=======
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessPID(pid())), nil
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Detect returns a *Resource that describes the name of the process executable.
func (processExecutableNameDetector) Detect(ctx context.Context) (*Resource, error) {
	executableName := filepath.Base(commandArgs()[0])

<<<<<<< HEAD
	return NewWithAttributes(semconv.ProcessExecutableNameKey.String(executableName)), nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessExecutableNameKey.String(executableName)), nil
=======
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessExecutableName(executableName)), nil
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Detect returns a *Resource that describes the full path of the process executable.
func (processExecutablePathDetector) Detect(ctx context.Context) (*Resource, error) {
	executablePath, err := executablePath()
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	return NewWithAttributes(semconv.ProcessExecutablePathKey.String(executablePath)), nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessExecutablePathKey.String(executablePath)), nil
=======
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessExecutablePath(executablePath)), nil
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Detect returns a *Resource that describes all the command arguments as received
// by the process.
func (processCommandArgsDetector) Detect(ctx context.Context) (*Resource, error) {
<<<<<<< HEAD
	return NewWithAttributes(semconv.ProcessCommandArgsKey.Array(commandArgs())), nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessCommandArgsKey.StringSlice(commandArgs())), nil
=======
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessCommandArgs(commandArgs()...)), nil
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Detect returns a *Resource that describes the username of the user that owns the
// process.
func (processOwnerDetector) Detect(ctx context.Context) (*Resource, error) {
	owner, err := owner()
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	return NewWithAttributes(semconv.ProcessOwnerKey.String(owner.Username)), nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessOwnerKey.String(owner.Username)), nil
=======
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessOwner(owner.Username)), nil
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Detect returns a *Resource that describes the name of the compiler used to compile
// this process image.
func (processRuntimeNameDetector) Detect(ctx context.Context) (*Resource, error) {
<<<<<<< HEAD
	return NewWithAttributes(semconv.ProcessRuntimeNameKey.String(runtimeName())), nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessRuntimeNameKey.String(runtimeName())), nil
=======
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessRuntimeName(runtimeName())), nil
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Detect returns a *Resource that describes the version of the runtime of this process.
func (processRuntimeVersionDetector) Detect(ctx context.Context) (*Resource, error) {
<<<<<<< HEAD
	return NewWithAttributes(semconv.ProcessRuntimeVersionKey.String(runtimeVersion())), nil
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessRuntimeVersionKey.String(runtimeVersion())), nil
=======
	return NewWithAttributes(semconv.SchemaURL, semconv.ProcessRuntimeVersion(runtimeVersion())), nil
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
}

// Detect returns a *Resource that describes the runtime of this process.
func (processRuntimeDescriptionDetector) Detect(ctx context.Context) (*Resource, error) {
	runtimeDescription := fmt.Sprintf(
		"go version %s %s/%s", runtimeVersion(), runtimeOS(), runtimeArch())

	return NewWithAttributes(
<<<<<<< HEAD
		semconv.ProcessRuntimeDescriptionKey.String(runtimeDescription),
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
		semconv.SchemaURL,
		semconv.ProcessRuntimeDescriptionKey.String(runtimeDescription),
=======
		semconv.SchemaURL,
		semconv.ProcessRuntimeDescription(runtimeDescription),
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	), nil
}

// WithProcessPID adds an attribute with the process identifier (PID) to the
// configured Resource.
func WithProcessPID() Option {
	return WithDetectors(processPIDDetector{})
}

// WithProcessExecutableName adds an attribute with the name of the process
// executable to the configured Resource.
func WithProcessExecutableName() Option {
	return WithDetectors(processExecutableNameDetector{})
}

// WithProcessExecutablePath adds an attribute with the full path to the process
// executable to the configured Resource.
func WithProcessExecutablePath() Option {
	return WithDetectors(processExecutablePathDetector{})
}

// WithProcessCommandArgs adds an attribute with all the command arguments (including
// the command/executable itself) as received by the process the configured Resource.
func WithProcessCommandArgs() Option {
	return WithDetectors(processCommandArgsDetector{})
}

// WithProcessOwner adds an attribute with the username of the user that owns the process
// to the configured Resource.
func WithProcessOwner() Option {
	return WithDetectors(processOwnerDetector{})
}

// WithProcessRuntimeName adds an attribute with the name of the runtime of this
// process to the configured Resource.
func WithProcessRuntimeName() Option {
	return WithDetectors(processRuntimeNameDetector{})
}

// WithProcessRuntimeVersion adds an attribute with the version of the runtime of
// this process to the configured Resource.
func WithProcessRuntimeVersion() Option {
	return WithDetectors(processRuntimeVersionDetector{})
}

// WithProcessRuntimeDescription adds an attribute with an additional description
// about the runtime of the process to the configured Resource.
func WithProcessRuntimeDescription() Option {
	return WithDetectors(processRuntimeDescriptionDetector{})
}

// WithProcess adds all the Process attributes to the configured Resource.
// See individual WithProcess* functions to configure specific attributes.
func WithProcess() Option {
	return WithDetectors(
		processPIDDetector{},
		processExecutableNameDetector{},
		processExecutablePathDetector{},
		processCommandArgsDetector{},
		processOwnerDetector{},
		processRuntimeNameDetector{},
		processRuntimeVersionDetector{},
		processRuntimeDescriptionDetector{},
	)
}
