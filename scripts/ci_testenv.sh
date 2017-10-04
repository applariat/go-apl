#!/bin/bash
export APL_API=https://api.applariat.io/v1/
export APL_SVC_USERNAME=email
export APL_SVC_PASSWORD=password
export APL_STACK_NAME="Acme Air Travel Application"
export APL_RELEASE_VERSION=1
export APL_COMPONENT_NAME="node"

export TRAVIS_COMMIT="517a29e"
export TRAVIS_TAG=""
export CREATE_RELEASE=true

. ./ci_deploy.sh
