#!/bin/bash

set -xe

# Testing Go client.
gotenberg --api-enable-basic-auth --webhook-client-timeout 10s &
sleep 10
export CGO_ENABLED=1
go test -race -cover -covermode=atomic github.com/dcaraxes/gotenberg-go-client/v8
sleep 10 # allows Gotenberg to remove generated files.