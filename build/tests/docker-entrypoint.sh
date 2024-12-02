#!/bin/bash

set -xe

# Testing Go client.
gotenberg --api-enable-basic-auth --chromium-auto-start --log-level error &
sleep 10
export CGO_ENABLED=1
go test -v -race -cover -covermode=atomic ./...
sleep 10 # allows Gotenberg to remove generated files.