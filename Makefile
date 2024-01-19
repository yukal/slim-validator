GitOwner = yukal
GitRepo = slim-validator
GitRepoLastCommit = $(shell git rev-list --tags --max-count=1)
GitRepoVer = $(shell git describe --tags $(GitRepoLastCommit))
CoverageIndicator = $(shell go test -cover -count=1 | grep coverage: | grep -Eo '[0-9]+\.[0-9]+')

default: lint

install-linter:
	@golangci-lint --version >/dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-linter
	golangci-lint run

fmt:
	gofmt -w .

vet:
	go vet

test:
# go test -race -coverprofile=coverage.out -count=15 .
	go test -coverpkg=. -coverprofile=coverage.out -count=15 . && \
	go tool cover -func=coverage.out

coverage: test
	go tool cover -html=coverage.out -o coverage.html

bench:
	go test -run Benchmark -bench=. -benchmem -benchtime=100000x .

prepare: fmt lint vet test
publish:
# https://go.dev/doc/modules/publishing
	GOPROXY=proxy.golang.org go list -m github.com/$(GitOwner)/$(GitRepo)@v$(GitRepoVer)

badge:
	wget --output-document=.github/badges/badge-coverage.svg https://badgen.net/badge/coverage/$(CoverageIndicator)%25/009000?labelColor=333
# wget --output-document=.github/badges/badge-license.svg https://badgen.net/github/license/$(GitOwner)/$(GitRepo)
# wget --output-document=.github/badges/badge-release.svg https://badgen.net/github/release/$(GitOwner)/$(GitRepo)

ver:
	@echo $(GitRepoVer)

.PHONY: lint fmt vet test coverage bench prepare publish badge ver install-linter
