#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

my_dir=$(dirname "${BASH_SOURCE}")


source ${my_dir}/../bin/common.sh

warn "assumes you have grpcc installed."

inf "grpcc --insecure --proto "${my_dir}/../api/hello.proto" --address ${SERVER_ADDRESS}  --exec ${my_dir}/../bin/verify_apis_script.js"
grpcc --insecure --proto "${my_dir}/../api/hello.proto" --address ${SERVER_ADDRESS}  --exec ${my_dir}/../bin/verify_apis_script.js
