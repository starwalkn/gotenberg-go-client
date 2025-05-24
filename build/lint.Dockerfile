ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION:-1.24.2}-alpine

ARG GOLANGCI_LINT_VERSION

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s | sh -s -- -b /usr/local/bin v${GOLANGCI_LINT_VERSION} &&\
    golangci-lint --version

# Define our workding outside of $GOPATH (we're using go modules).
WORKDIR /lint

# Copy our module dependencies definitions.
COPY go.mod .
COPY go.sum .

# Copy golangci-lint configuration file.
COPY .golangci.yml .

# Install module dependencies.
RUN go mod download

CMD [ "golangci-lint", "run"]