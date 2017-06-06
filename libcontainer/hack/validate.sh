#!/usr/bin/env bash
set -e

# This script runs all validations

validate() {
    export MAKEDIR=/go/src/github.com/nicle-lin/dockerV1.2.0/docker/hack/make
    sed -i 's!docker/docker!docker/libcontainer!' /go/src/github.com/nicle-lin/dockerV1.2.0/docker/hack/make/.validate
    bash /go/src/github.com/nicle-lin/dockerV1.2.0/docker/hack/make/validate-dco
    bash /go/src/github.com/nicle-lin/dockerV1.2.0/docker/hack/make/validate-gofmt
    go get golang.org/x/tools/cmd/vet
    bash /go/src/github.com/nicle-lin/dockerV1.2.0/docker/hack/make/validate-vet
}

# run validations
validate
