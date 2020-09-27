GIT_BIN := git

GO_BIN := go
GO_FMT_BIN := gofmt
GO_LINT := $(GO_BIN) run ./vendor/golang.org/x/lint/golint
GO_STATICCHECK_BIN := $(GO_BIN) run ./vendor/honnef.co/go/tools/cmd/staticcheck

test: test-style

test-style: test-fmt test-lint test-tidy test-staticcheck

test-fmt:
	@echo "+ $@"
	@test -z "$$($(GO_FMT_BIN) -l -e -s . | grep /openapi-clients/ | tee /dev/stderr)" || \
	  ( >&2 echo "=> please format Go code with '$(GO_FMT_BIN) -s -w .'" && false)

test-lint:
	@echo "+ $@"
	@test -z "$$($(GO_LINT) main.go | tee /dev/stderr )"

test-tidy:
	@echo "+ $@"
	@$(GO_BIN) mod tidy
	@test -z "$$($(GIT_BIN) status --short go.mod go.sum | tee /dev/stderr)" || \
	  ( >&2 echo "=> please tidy the Go modules with '$(GO_BIN) mod tidy'" && false)

test-staticcheck:
	@echo "+ $@"
	@$(GO_STATICCHECK_BIN) ./...
