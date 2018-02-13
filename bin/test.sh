#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

ROOT=$(dirname "${BASH_SOURCE}")/..
source "${ROOT}/bin/common.sh"

while getopts ":v" opt; do
  case ${opt} in
    v)
      SILENT=false
      ;;
    \?)
      echo "Invalid flag: -${OPTARG}" >&2
      exit 1
      ;;
  esac
done

if ${SILENT} ; then
  echo "Running in silent mode, run with -v if you want to see script logs."
fi

EXIT_CODE=0

echo
inf "running static code checks..."
inf "${ROOT}/bin/verify.sh"
run-cmd bash "${ROOT}/bin/verify.sh" && tr=$? || tr=$?
if [[ ! ${tr} -eq 0 ]]; then
  EXIT_CODE=1
fi

FAILED_TESTS=()

echo
inf "run go tests on packages..."
for pkg in $(packages); do
  go test ${pkg} -coverprofile=${pkg%/}-coverage.out && tr=$? || tr=$?
  inf "coverage, if created, can be found in ${pkg%/}-coverage.out"
  if [[ ! ${tr} -eq 0 ]]; then
    FAILED_TESTS+="${pkg}"
    EXIT_CODE=1
  fi
done

if [[ ${EXIT_CODE} -eq 1 ]]; then
    print-failed-go-test-package
fi

exit ${EXIT_CODE}

