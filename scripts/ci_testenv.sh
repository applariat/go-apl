#!/bin/bash
#CLI Version to download
export APL_CLI_VER=v0.2.0

#Project variables
export REPO_PATH="https://github.com/applariat/acme-air/archive"

#APL Platform variables
#Required as env variable inputs from CI
export APL_LOC_DEPLOY_NAME="w3-tutorial"
export APL_LOC_ARTIFACT_NAME="simple-url"

#APL STACK variables
#Required as env variable inputs from CI
export APL_STACK_NAME="acmeair"
export APL_RELEASE_VERSION=1
export APL_COMPONENT_NAME="acmeair"
export APL_ARTIFACT_NAME="acme-air"

export TRAVIS_COMMIT="517a29e"
export TRAVIS_TAG="test3"

. ./ci_deploy.sh
