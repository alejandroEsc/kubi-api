#!/bin/bash

# from http://github.com/kubernetes/kubernetes/hack/verify-gofmt.sh

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(dirname "${BASH_SOURCE}")/../..

cd "${ROOT}"

goconst=$(which goconst)
if [[ ! -x "${goconst}" ]]; then
  warn "could not find goconst, please verify your GOPATH"
  inf "https://github.com/jgautheron/goconst"
  exit 1
fi

source "${ROOT}/bin/common.sh"

diff=$( echo `packages` | xargs ${goconst} -min-occurrences 2 2>&1) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}"
  exit 1
fi
