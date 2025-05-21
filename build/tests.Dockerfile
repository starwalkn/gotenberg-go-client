ARG GOLANG_VERSION
ARG GOTENBERG_VERSION

FROM golang:${GOLANG_VERSION:-1.24.2}-alpine AS golang

FROM gotenberg/gotenberg:${GOTENBERG_VERSION:-8.20.1}

USER root

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN apt-get update \
  && apt-get install -y --no-install-recommends build-essential \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

COPY --from=golang /usr/local/go /usr/local/go

RUN go version

ENV GOTENBERG_API_BASIC_AUTH_USERNAME=foo
ENV GOTENBERG_API_BASIC_AUTH_PASSWORD=bar

# Define our workding outside of $GOPATH (we're using go modules).
WORKDIR /tests

# Copy our module dependencies definitions.
COPY go.mod .
COPY go.sum .

# Install module dependencies.
RUN go mod download

USER gotenberg

ENTRYPOINT [ "build/tests-entrypoint.sh" ]
