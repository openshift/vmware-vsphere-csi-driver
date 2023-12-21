# Copyright The OpenTelemetry Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

EXAMPLES := $(shell ./get_main_pkgs.sh ./example)
TOOLS_MOD_DIR := ./internal/tools

# All source code and documents. Used in spell check.
ALL_DOCS := $(shell find . -name '*.md' -type f | sort)
<<<<<<< HEAD
# All directories with go.mod files related to opentelemetry library. Used for building, testing and linting.
ALL_GO_MOD_DIRS := $(filter-out $(TOOLS_MOD_DIR), $(shell find . -type f -name 'go.mod' -exec dirname {} \; | egrep -v '^./example' | sort)) $(shell find ./example -type f -name 'go.mod' -exec dirname {} \; | sort)
ALL_COVERAGE_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | egrep -v '^./example|^$(TOOLS_MOD_DIR)' | sort)
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
ALL_GO_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | sort)
OTEL_GO_MOD_DIRS := $(filter-out $(TOOLS_MOD_DIR), $(ALL_GO_MOD_DIRS))
ALL_COVERAGE_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | egrep -v '^./example|^$(TOOLS_MOD_DIR)' | sort)
=======
ALL_GO_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | sort)
OTEL_GO_MOD_DIRS := $(filter-out $(TOOLS_MOD_DIR), $(ALL_GO_MOD_DIRS))
ALL_COVERAGE_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | grep -E -v '^./example|^$(TOOLS_MOD_DIR)' | sort)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

GO = go
TIMEOUT = 60

.DEFAULT_GOAL := precommit

.PHONY: precommit ci
<<<<<<< HEAD
precommit: dependabot-check license-check lint build examples test-default
ci: precommit check-clean-work-tree test-coverage
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
precommit: dependabot-generate license-check vanity-import-fix misspell go-mod-tidy golangci-lint-fix test-default
ci: dependabot-check license-check lint vanity-import-check build test-default check-clean-work-tree test-coverage
=======
precommit: generate dependabot-generate license-check misspell go-mod-tidy golangci-lint-fix test-default
ci: generate dependabot-check license-check lint vanity-import-check build test-default check-clean-work-tree test-coverage
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

# Tools

TOOLS = $(CURDIR)/.tools

$(TOOLS):
	@mkdir -p $@
$(TOOLS)/%: | $(TOOLS)
	cd $(TOOLS_MOD_DIR) && \
	$(GO) build -o $@ $(PACKAGE)

CROSSLINK = $(TOOLS)/crosslink
$(TOOLS)/crosslink: PACKAGE=go.opentelemetry.io/otel/$(TOOLS_MOD_DIR)/crosslink

GOLANGCI_LINT = $(TOOLS)/golangci-lint
$(TOOLS)/golangci-lint: PACKAGE=github.com/golangci/golangci-lint/cmd/golangci-lint

MISSPELL = $(TOOLS)/misspell
$(TOOLS)/misspell: PACKAGE= github.com/client9/misspell/cmd/misspell

STRINGER = $(TOOLS)/stringer
$(TOOLS)/stringer: PACKAGE=golang.org/x/tools/cmd/stringer

$(TOOLS)/gojq: PACKAGE=github.com/itchyny/gojq/cmd/gojq

GOTMPL = $(TOOLS)/gotmpl
$(GOTMPL): PACKAGE=go.opentelemetry.io/build-tools/gotmpl

GORELEASE = $(TOOLS)/gorelease
$(GORELEASE): PACKAGE=golang.org/x/exp/cmd/gorelease

GOVULNCHECK = $(TOOLS)/govulncheck
$(TOOLS)/govulncheck: PACKAGE=golang.org/x/vuln/cmd/govulncheck

.PHONY: tools
<<<<<<< HEAD
tools: $(CROSSLINK) $(GOLANGCI_LINT) $(MISSPELL) $(STRINGER) $(TOOLS)/gojq

||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
tools: $(CROSSLINK) $(DBOTCONF) $(GOLANGCI_LINT) $(MISSPELL) $(GOCOVMERGE) $(STRINGER) $(PORTO) $(GOJQ) $(SEMCONVGEN) $(MULTIMOD) $(SEMCONVKIT)
=======
tools: $(CROSSLINK) $(DBOTCONF) $(GOLANGCI_LINT) $(MISSPELL) $(GOCOVMERGE) $(STRINGER) $(PORTO) $(GOJQ) $(SEMCONVGEN) $(MULTIMOD) $(SEMCONVKIT) $(GOTMPL) $(GORELEASE)

# Virtualized python tools via docker

# The directory where the virtual environment is created.
VENVDIR := venv

# The directory where the python tools are installed.
PYTOOLS := $(VENVDIR)/bin

# The pip executable in the virtual environment.
PIP := $(PYTOOLS)/pip

# The directory in the docker image where the current directory is mounted.
WORKDIR := /workdir

# The python image to use for the virtual environment.
PYTHONIMAGE := python:3.11.3-slim-bullseye

# Run the python image with the current directory mounted.
DOCKERPY := docker run --rm -v "$(CURDIR):$(WORKDIR)" -w $(WORKDIR) $(PYTHONIMAGE)

# Create a virtual environment for Python tools.
$(PYTOOLS):
# The `--upgrade` flag is needed to ensure that the virtual environment is
# created with the latest pip version.
	@$(DOCKERPY) bash -c "python3 -m venv $(VENVDIR) && $(PIP) install --upgrade pip"

# Install python packages into the virtual environment.
$(PYTOOLS)/%: | $(PYTOOLS)
	@$(DOCKERPY) $(PIP) install -r requirements.txt

CODESPELL = $(PYTOOLS)/codespell
$(CODESPELL): PACKAGE=codespell

# Generate

.PHONY: generate
generate: go-generate vanity-import-fix

.PHONY: go-generate
go-generate: $(OTEL_GO_MOD_DIRS:%=go-generate/%)
go-generate/%: DIR=$*
go-generate/%: | $(STRINGER) $(GOTMPL)
	@echo "$(GO) generate $(DIR)/..." \
		&& cd $(DIR) \
		&& PATH="$(TOOLS):$${PATH}" $(GO) generate ./...

.PHONY: vanity-import-fix
vanity-import-fix: | $(PORTO)
	@$(PORTO) --include-internal -w .

# Generate go.work file for local development.
.PHONY: go-work
go-work: | $(CROSSLINK)
	$(CROSSLINK) work --root=$(shell pwd)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

# Build

<<<<<<< HEAD
.PHONY: examples generate build
examples:
	@set -e; for dir in $(EXAMPLES); do \
	  echo "$(GO) build $${dir}/..."; \
	  (cd "$${dir}" && \
	   $(GO) build .); \
	done
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
.PHONY: generate build
=======
.PHONY: build
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

<<<<<<< HEAD
generate: $(STRINGER)
	set -e; for dir in $(ALL_GO_MOD_DIRS); do \
	  echo "$(GO) generate $${dir}/..."; \
	  (cd "$${dir}" && \
	    PATH="$(TOOLS):$${PATH}" $(GO) generate ./...); \
	done

build: generate
	# Build all package code including testing code.
	set -e; for dir in $(ALL_GO_MOD_DIRS); do \
	  echo "$(GO) build $${dir}/..."; \
	  (cd "$${dir}" && \
	    $(GO) build ./... && \
		$(GO) list ./... \
		  | grep -v third_party \
		  | xargs $(GO) test -vet=off -run xxxxxMatchNothingxxxxx >/dev/null); \
	done
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
generate: $(OTEL_GO_MOD_DIRS:%=generate/%)
generate/%: DIR=$*
generate/%: | $(STRINGER) $(PORTO)
	@echo "$(GO) generate $(DIR)/..." \
		&& cd $(DIR) \
		&& PATH="$(TOOLS):$${PATH}" $(GO) generate ./... && $(PORTO) -w .

build: generate $(OTEL_GO_MOD_DIRS:%=build/%) $(OTEL_GO_MOD_DIRS:%=build-tests/%)
build/%: DIR=$*
build/%:
	@echo "$(GO) build $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) build ./...

build-tests/%: DIR=$*
build-tests/%:
	@echo "$(GO) build tests $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) list ./... \
		| grep -v third_party \
		| xargs $(GO) test -vet=off -run xxxxxMatchNothingxxxxx >/dev/null
=======
build: $(OTEL_GO_MOD_DIRS:%=build/%) $(OTEL_GO_MOD_DIRS:%=build-tests/%)
build/%: DIR=$*
build/%:
	@echo "$(GO) build $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) build ./...

build-tests/%: DIR=$*
build-tests/%:
	@echo "$(GO) build tests $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) list ./... \
		| grep -v third_party \
		| xargs $(GO) test -vet=off -run xxxxxMatchNothingxxxxx >/dev/null
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test
test-default: ARGS=-v -race
test-bench:   ARGS=-run=xxxxxMatchNothingxxxxx -test.benchtime=1ms -bench=.
test-short:   ARGS=-short
test-verbose: ARGS=-v
test-race:    ARGS=-race
$(TEST_TARGETS): test
test:
	@set -e; for dir in $(ALL_GO_MOD_DIRS); do \
	  echo "$(GO) test -timeout $(TIMEOUT)s $(ARGS) $${dir}/..."; \
	  (cd "$${dir}" && \
	    $(GO) list ./... \
		  | grep -v third_party \
		  | xargs $(GO) test -timeout $(TIMEOUT)s $(ARGS)); \
	done

COVERAGE_MODE    = atomic
COVERAGE_PROFILE = coverage.out
.PHONY: test-coverage
test-coverage:
	@set -e; \
	printf "" > coverage.txt; \
	for dir in $(ALL_COVERAGE_MOD_DIRS); do \
	  echo "$(GO) test -coverpkg=./... -covermode=$(COVERAGE_MODE) -coverprofile="$(COVERAGE_PROFILE)" $${dir}/..."; \
	  (cd "$${dir}" && \
	    $(GO) list ./... \
	    | grep -v third_party \
	    | xargs $(GO) test -coverpkg=./... -covermode=$(COVERAGE_MODE) -coverprofile="$(COVERAGE_PROFILE)" && \
	  $(GO) tool cover -html=coverage.out -o coverage.html); \
	  [ -f "$${dir}/coverage.out" ] && cat "$${dir}/coverage.out" >> coverage.txt; \
	done; \
<<<<<<< HEAD
	sed -i.bak -e '2,$$ { /^mode: /d; }' coverage.txt
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	$(GOCOVMERGE) $$(find . -name coverage.out) > coverage.txt

.PHONY: golangci-lint golangci-lint-fix
golangci-lint-fix: ARGS=--fix
golangci-lint-fix: golangci-lint
golangci-lint: $(OTEL_GO_MOD_DIRS:%=golangci-lint/%)
golangci-lint/%: DIR=$*
golangci-lint/%: | $(GOLANGCI_LINT)
	@echo 'golangci-lint $(if $(ARGS),$(ARGS) ,)$(DIR)' \
		&& cd $(DIR) \
		&& $(GOLANGCI_LINT) run --allow-serial-runners $(ARGS)

.PHONY: crosslink
crosslink: | $(CROSSLINK)
	@echo "Updating intra-repository dependencies in all go modules" \
		&& $(CROSSLINK) --root=$(shell pwd) --prune

.PHONY: go-mod-tidy
go-mod-tidy: $(ALL_GO_MOD_DIRS:%=go-mod-tidy/%)
go-mod-tidy/%: DIR=$*
go-mod-tidy/%: | crosslink
	@echo "$(GO) mod tidy in $(DIR)" \
		&& cd $(DIR) \
		&& $(GO) mod tidy -compat=1.17

.PHONY: lint-modules
lint-modules: go-mod-tidy
=======
	$(GOCOVMERGE) $$(find . -name coverage.out) > coverage.txt

# Adding a directory will include all benchmarks in that direcotry if a filter is not specified.
BENCHMARK_TARGETS := sdk/trace
.PHONY: benchmark
benchmark: $(BENCHMARK_TARGETS:%=benchmark/%)
BENCHMARK_FILTER = .
# You can override the filter for a particular directory by adding a rule here.
benchmark/sdk/trace: BENCHMARK_FILTER = SpanWithAttributes_8/AlwaysSample
benchmark/%:
	@echo "$(GO) test -timeout $(TIMEOUT)s -run=xxxxxMatchNothingxxxxx -bench=$(BENCHMARK_FILTER) $*..." \
		&& cd $* \
		$(foreach filter, $(BENCHMARK_FILTER), && $(GO) test -timeout $(TIMEOUT)s -run=xxxxxMatchNothingxxxxx -bench=$(filter))

.PHONY: golangci-lint golangci-lint-fix
golangci-lint-fix: ARGS=--fix
golangci-lint-fix: golangci-lint
golangci-lint: $(OTEL_GO_MOD_DIRS:%=golangci-lint/%)
golangci-lint/%: DIR=$*
golangci-lint/%: | $(GOLANGCI_LINT)
	@echo 'golangci-lint $(if $(ARGS),$(ARGS) ,)$(DIR)' \
		&& cd $(DIR) \
		&& $(GOLANGCI_LINT) run --allow-serial-runners $(ARGS)

.PHONY: crosslink
crosslink: | $(CROSSLINK)
	@echo "Updating intra-repository dependencies in all go modules" \
		&& $(CROSSLINK) --root=$(shell pwd) --prune

.PHONY: go-mod-tidy
go-mod-tidy: $(ALL_GO_MOD_DIRS:%=go-mod-tidy/%)
go-mod-tidy/%: DIR=$*
go-mod-tidy/%: | crosslink
	@echo "$(GO) mod tidy in $(DIR)" \
		&& cd $(DIR) \
		&& $(GO) mod tidy -compat=1.20

.PHONY: lint-modules
lint-modules: go-mod-tidy
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

.PHONY: lint
<<<<<<< HEAD
lint: misspell lint-modules | $(GOLANGCI_LINT)
	set -e; for dir in $(ALL_GO_MOD_DIRS); do \
	  echo "golangci-lint in $${dir}"; \
	  (cd "$${dir}" && \
	    $(GOLANGCI_LINT) run --fix && \
	    $(GOLANGCI_LINT) run); \
	done
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
lint: misspell lint-modules golangci-lint

.PHONY: vanity-import-check
vanity-import-check: | $(PORTO)
	@$(PORTO) --include-internal -l . || echo "(run: make vanity-import-fix)"

.PHONY: vanity-import-fix
vanity-import-fix: | $(PORTO)
	@$(PORTO) --include-internal -w .
=======
lint: misspell lint-modules golangci-lint govulncheck

.PHONY: vanity-import-check
vanity-import-check: | $(PORTO)
	@$(PORTO) --include-internal -l . || ( echo "(run: make vanity-import-fix)"; exit 1 )
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

.PHONY: misspell
misspell: | $(MISSPELL)
	$(MISSPELL) -w $(ALL_DOCS)

.PHONY: lint-modules
lint-modules: | $(CROSSLINK)
	set -e; for dir in $(ALL_GO_MOD_DIRS) $(TOOLS_MOD_DIR); do \
	  echo "$(GO) mod tidy in $${dir}"; \
	  (cd "$${dir}" && \
	    $(GO) mod tidy); \
	done
	echo "cross-linking all go modules"
	$(CROSSLINK)

.PHONY: govulncheck
govulncheck: $(OTEL_GO_MOD_DIRS:%=govulncheck/%)
govulncheck/%: DIR=$*
govulncheck/%: | $(GOVULNCHECK)
	@echo "govulncheck ./... in $(DIR)" \
		&& cd $(DIR) \
		&& $(GOVULNCHECK) ./...

.PHONY: codespell
codespell: | $(CODESPELL)
	@$(DOCKERPY) $(CODESPELL)

.PHONY: license-check
license-check:
<<<<<<< HEAD
	@licRes=$$(for f in $$(find . -type f \( -iname '*.go' -o -iname '*.sh' \) ! -path '**/third_party/*' ! -path './exporters/otlp/internal/opentelemetry-proto/*') ; do \
	           awk '/Copyright The OpenTelemetry Authors|generated|GENERATED/ && NR<=3 { found=1; next } END { if (!found) print FILENAME }' $$f; \
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	@licRes=$$(for f in $$(find . -type f \( -iname '*.go' -o -iname '*.sh' \) ! -path '**/third_party/*' ! -path './.git/*' ) ; do \
	           awk '/Copyright The OpenTelemetry Authors|generated|GENERATED/ && NR<=3 { found=1; next } END { if (!found) print FILENAME }' $$f; \
=======
	@licRes=$$(for f in $$(find . -type f \( -iname '*.go' -o -iname '*.sh' \) ! -path '**/third_party/*' ! -path './.git/*' ) ; do \
	           awk '/Copyright The OpenTelemetry Authors|generated|GENERATED/ && NR<=4 { found=1; next } END { if (!found) print FILENAME }' $$f; \
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
	   done); \
	   if [ -n "$${licRes}" ]; then \
	           echo "license header checking failed:"; echo "$${licRes}"; \
	           exit 1; \
	   fi

.PHONY: dependabot-check
<<<<<<< HEAD
dependabot-check:
	@result=$$( \
		for f in $$( find . -type f -name go.mod -exec dirname {} \; | sed 's/^.\/\?/\//' ); \
			do grep -q "$$f" .github/dependabot.yml \
			|| echo "$$f"; \
		done; \
	); \
	if [ -n "$$result" ]; then \
		echo "missing go.mod dependabot check:"; echo "$$result"; \
		exit 1; \
	fi
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
dependabot-check: | $(DBOTCONF)
	@$(DBOTCONF) verify $(DEPENDABOT_CONFIG) || echo "(run: make dependabot-generate)"

.PHONY: dependabot-generate
dependabot-generate: | $(DBOTCONF)
	@$(DBOTCONF) generate > $(DEPENDABOT_CONFIG)
=======
dependabot-check: | $(DBOTCONF)
	@$(DBOTCONF) verify $(DEPENDABOT_CONFIG) || ( echo "(run: make dependabot-generate)"; exit 1 )

.PHONY: dependabot-generate
dependabot-generate: | $(DBOTCONF)
	@$(DBOTCONF) generate > $(DEPENDABOT_CONFIG)
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

.PHONY: check-clean-work-tree
check-clean-work-tree:
	@if ! git diff --quiet; then \
	  echo; \
	  echo 'Working tree is not clean, did you forget to run "make precommit"?'; \
	  echo; \
	  git status; \
	  exit 1; \
	fi
<<<<<<< HEAD
||||||| parent of 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))

SEMCONVPKG ?= "semconv/"
.PHONY: semconv-generate
semconv-generate: | $(SEMCONVGEN) $(SEMCONVKIT)
	@[ "$(TAG)" ] || ( echo "TAG unset: missing opentelemetry specification tag"; exit 1 )
	@[ "$(OTEL_SPEC_REPO)" ] || ( echo "OTEL_SPEC_REPO unset: missing path to opentelemetry specification repo"; exit 1 )
	@$(SEMCONVGEN) -i "$(OTEL_SPEC_REPO)/semantic_conventions/trace" -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	@$(SEMCONVGEN) -i "$(OTEL_SPEC_REPO)/semantic_conventions/resource" -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	@$(SEMCONVKIT) -output "$(SEMCONVPKG)/$(TAG)" -tag "$(TAG)"

.PHONY: prerelease
prerelease: | $(MULTIMOD)
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) verify && $(MULTIMOD) prerelease -m ${MODSET}

COMMIT ?= "HEAD"
.PHONY: add-tags
add-tags: | $(MULTIMOD)
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) verify && $(MULTIMOD) tag -m ${MODSET} -c ${COMMIT}
=======

SEMCONVPKG ?= "semconv/"
.PHONY: semconv-generate
semconv-generate: | $(SEMCONVGEN) $(SEMCONVKIT)
	[ "$(TAG)" ] || ( echo "TAG unset: missing opentelemetry semantic-conventions tag"; exit 1 )
	[ "$(OTEL_SEMCONV_REPO)" ] || ( echo "OTEL_SEMCONV_REPO unset: missing path to opentelemetry semantic-conventions repo"; exit 1 )
	$(SEMCONVGEN) -i "$(OTEL_SEMCONV_REPO)/model/." --only=span -p conventionType=trace -f trace.go -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	$(SEMCONVGEN) -i "$(OTEL_SEMCONV_REPO)/model/." --only=attribute_group -p conventionType=trace -f attribute_group.go -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	$(SEMCONVGEN) -i "$(OTEL_SEMCONV_REPO)/model/." --only=event -p conventionType=event -f event.go -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	$(SEMCONVGEN) -i "$(OTEL_SEMCONV_REPO)/model/." --only=resource -p conventionType=resource -f resource.go -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	$(SEMCONVKIT) -output "$(SEMCONVPKG)/$(TAG)" -tag "$(TAG)"

.PHONY: gorelease
gorelease: $(OTEL_GO_MOD_DIRS:%=gorelease/%)
gorelease/%: DIR=$*
gorelease/%:| $(GORELEASE)
	@echo "gorelease in $(DIR):" \
		&& cd $(DIR) \
		&& $(GORELEASE) \
		|| echo ""

.PHONY: prerelease
prerelease: | $(MULTIMOD)
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) verify && $(MULTIMOD) prerelease -m ${MODSET}

COMMIT ?= "HEAD"
.PHONY: add-tags
add-tags: | $(MULTIMOD)
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) verify && $(MULTIMOD) tag -m ${MODSET} -c ${COMMIT}

.PHONY: lint-markdown
lint-markdown: 
	docker run -v "$(CURDIR):$(WORKDIR)" docker://avtodev/markdown-lint:v1 -c $(WORKDIR)/.markdownlint.yaml $(WORKDIR)/**/*.md
>>>>>>> 60945b63 (UPSTREAM: 2686: Bump OpenTelemetry libs (#2686))
