#!/bin/bash
export APL_STACK_NAME="acme-air"
export APL_RELEASE_VERSION=1
export APL_COMPONENT_NAME="node"

export TRAVIS_COMMIT="517a29e"
export TRAVIS_TAG="test3"

. ./ci_deploy.sh
