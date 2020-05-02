#!/bin/sh

set -e

for CMD in `ls cmd`; do
  go install -v ./cmd/$CMD
done
