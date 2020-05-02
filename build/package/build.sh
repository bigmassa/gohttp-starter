#!/bin/sh

set -e

for CMD in `ls cmd`; do
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $CMD ./cmd/$CMD
done
