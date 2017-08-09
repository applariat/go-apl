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

BIN_NAME=apl-${VERSION}-${OS}_${ARCH}

go build -ldflags "-X github.com/applariat/go-apl/cmd/apl/app.VERSION=${VERSION}" -o bin/apl cmd/apl/main.go

if [[ ${OS} == windows ]]; then
	tar -czf bin/${BIN_NAME}.tgz -C bin apl
else
	tar -czf bin/${BIN_NAME}.tgz bin/apl scripts apl_install.sh
fi
