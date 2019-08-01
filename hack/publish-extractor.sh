#!/usr/bin/env sh

set -eux
cd "${0%/*}"/..

KO_DOCKER_REPO=gcr.io/kf-releases ko publish github.com/poy/kontext/cmd/extractor -P
