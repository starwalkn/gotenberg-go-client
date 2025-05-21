ARG GOLANG_VERSION
ARG GOTENBERG_VERSION

FROM golang:${GOLANG_VERSION:-1.24.2}-alpine AS golang

FROM gotenberg/gotenberg:${GOTENBERG_VERSION:-8.20.1}

USER root

# |--------------------------------------------------------------------------
# | Common libraries
# |--------------------------------------------------------------------------
# |
# | Libraries used in the build process of this image.
# |
RUN apt-get update \
  && apt-get install -y build-essential \
  && apt-get install manpages-dev

# |--------------------------------------------------------------------------
# | Golang
# |--------------------------------------------------------------------------
# |
# | Installs Golang.
# |

COPY --from=golang /usr/local/go /usr/local/go

RUN export PATH="/usr/local/go/bin:$PATH" &&\
    go version

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# |--------------------------------------------------------------------------
# | Final touch
# |--------------------------------------------------------------------------
# |
# | Last instructions of this build.
# |

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
