#!/bin/bash

set -e

# Run Gotenberg in the background.
gotenberg --api-enable-basic-auth --chromium-auto-start --log-level error &

GOTENBERG_HEALTH_URL="http://localhost:3000/health"

# Waiting for Gotenberg to be ready.
echo "Waiting for Gotenberg to be ready..."
MAX_ATTEMPTS=30
ATTEMPTS=0

while [ "$ATTEMPTS" -lt "$MAX_ATTEMPTS" ]; do
  if curl --fail --silent -o /dev/null "$GOTENBERG_HEALTH_URL"; then
    echo "Gotenberg is ready."
    break
  fi
  echo "Gotenberg not ready yet, retrying in 1 second..."
  sleep 1
  ATTEMPTS=$((ATTEMPTS+1))
done

# If Gotenberg does not start within the allotted time, terminate the script with an error.
if [ "$ATTEMPTS" -eq "$MAX_ATTEMPTS" ]; then
  echo "Gotenberg failed to start within the expected time."
  exit 1
fi

echo "Running Go tests..."
go test -v -race -cover -covermode=atomic ./...
