#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [ -z "${OS}" ]; then
    echo "OS must be set"
    exit 1
fi

if [ -z "${ARCH}" ]; then
    echo "ARCH must be set"
    exit 1
fi

if [ -z "${VERSION}" ]; then
    echo "VERSION must be set"
    exit 1
fi

export CGO_ENABLED=0
export GOARCH="${ARCH}"
export GOOS="${OS}"

BIN_NAME=bin/apl-${VERSION}-${OS}_${ARCH}

go build -ldflags "-X github.com/applariat/go-apl/cmd/apl/app.VERSION=${VERSION}" -o ${BIN_NAME} cmd/apl/main.go

tar -czf ${BIN_NAME}.tgz ${BIN_NAME}
