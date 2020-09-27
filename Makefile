BUILD_FOLDER = dist
BUILD_OPTIONS =

OS := darwin
ARCH := amd64
BINARY_NAME := "b64"
FULL_BINARY_NAME := $(BINARY_NAME)-$(OS)-$(ARCH)

PROJECT_USERNAME := kdisneur
PROJECT_REPOSITORY := b64

GIT_BIN := git

GIT_TAG := $(shell $(GIT_BIN) tag --points-at HEAD)
GIT_COMMIT := $(shell $(GIT_BIN) rev-parse HEAD)
GIT_BRANCH := $(shell $(GIT_BIN) branch --no-color | awk '/^\* / { print $$2 }')
GIT_STATE := $(shell if [ -z "$(shell $(GIT_BIN) status --short)" ]; then echo clean; else echo dirty; fi)
ALREADY_RELEASED := $(shell if [ $$(curl --silent --output /dev/null --write-out "%{http_code}" https://api.github.com/repos/$(PROJECT_USERNAME)/$(PROJECT_REPOSITORY)/releases/tags/$(GIT_TAG)) -eq 200 ]; then echo "true"; else echo "false"; fi)
PRERELASE := $(shell if [ "$(GIT_TAG)" != "" ]; then echo "true"; else echo "false"; fi)

GO_BIN := go
GO_FMT_BIN := gofmt
GO_LINT_BIN := $(GO_BIN) run ./vendor/golang.org/x/lint/golint
GO_STATICCHECK_BIN := $(GO_BIN) run ./vendor/honnef.co/go/tools/cmd/staticcheck

compile:
	@touch internal/version.go

	GOOS=$(OS) GOARCH=$(ARCH) go build $(BUILD_OPTIONS) \
		-ldflags \
		"-X github.com/kdisneur/b64/internal.versionNumber=$$(if [ "$(GIT_TAG)" = "" ]; then echo "unknown"; else echo "$(GIT_TAG)"; fi) \
			 -X github.com/kdisneur/b64/internal.gitBranch=$(GIT_BRANCH) \
			 -X github.com/kdisneur/b64/internal.gitCommit=$(GIT_COMMIT) \
			 -X github.com/kdisneur/b64/internal.gitState=$(GIT_STATE) \
			 -X github.com/kdisneur/b64/internal.buildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')" \
		-o $(BUILD_FOLDER)/$(BINARY_NAME)

	@tar czf $(BUILD_FOLDER)/$(FULL_BINARY_NAME).tgz -C $(BUILD_FOLDER) $(BINARY_NAME)
	@echo "archive generated at $(BUILD_FOLDER)/$(FULL_BINARY_NAME).tgz"

	@mv $(BUILD_FOLDER)/$(BINARY_NAME) $(BUILD_FOLDER)/$(FULL_BINARY_NAME)
	@echo "archive generated at $(BUILD_FOLDER)/$(FULL_BINARY_NAME)"

test: test-style

test-style: test-fmt test-lint test-tidy test-staticcheck

test-fmt:
	@echo "+ $@"
	@test -z "$$($(GO_FMT_BIN) -l -e -s . | tee /dev/stderr)" || \
	  ( >&2 echo "=> please format Go code with '$(GO_FMT_BIN) -s -w .'" && false)

test-lint:
	@echo "+ $@"
	@test -z "$$($(GO_LINT_BIN) internal . | tee /dev/stderr )"

test-tidy:
	@echo "+ $@"
	@$(GO_BIN) mod tidy
	@test -z "$$($(GIT_BIN) status --short go.mod go.sum | tee /dev/stderr)" || \
	  ( >&2 echo "=> please tidy the Go modules with '$(GO_BIN) mod tidy'" && false)

test-staticcheck:
	@echo "+ $@"
	@$(GO_STATICCHECK_BIN) ./...
