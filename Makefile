GOLANG_VERSION=1.23.2
GOTENBERG_VERSION=edge
GOLANGCI_LINT_VERSION=1.61.0

REPO=starwalkn/gotenberg-go-client/v8

# gofumpt and goimports all go files.
fmt:
	gofumpt -l -w .
	go mod tidy

# run linters.
lint:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) --build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION) -t $(REPO):lint -f build/lint/Dockerfile .
	docker run --rm -t -v "$(PWD):/lint" $(REPO):lint

# run all tests.
tests:
	docker build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) --build-arg GOTENBERG_VERSION=$(GOTENBERG_VERSION) -t $(REPO):tests -f build/tests/Dockerfile .
	docker run --rm -t -v "$(PWD):/tests" $(REPO):tests