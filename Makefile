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
	go test -coverpkg=. -coverprofile=coverage.out -count=15 . && \
	go tool cover -func=coverage.out

coverage:
# @go test -race -coverprofile=coverage.out .
	go test -coverpkg=. -coverprofile=coverage.out -count=15 . && \
	go tool cover -html=coverage.out -o coverage.html

bench:
	go test -run Benchmark -bench=. -benchmem -benchtime=100000x .

publish: fmt lint vet test

.PHONY: lint fmt vet test coverage bench install-linter publish
