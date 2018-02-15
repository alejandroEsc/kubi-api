#!/bin/bash

# from http://github.com/kubernetes/kubernetes/hack/verify-gofmt.sh

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(dirname "${BASH_SOURCE}")/../..

cd "${ROOT}"

gas=$(which gas)
if [[ ! -x "${gas}" ]]; then
  warn "could not find goconst, please verify your GOPATH"
  inf "https://github.com/GoASTScanner/gas"
  exit 1
fi

source "${ROOT}/bin/common.sh"

inf "false-positives may be mitigated, please read https://github.com/GoASTScanner/gas for details."
gas --exclude G204 ./... 2>&1
