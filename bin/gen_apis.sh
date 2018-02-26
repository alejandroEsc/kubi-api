#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

my_dir=$(dirname "${BASH_SOURCE}")

source ${my_dir}/../bin/common.sh

echo
inf "generating api stubs..."
inf "protoc -I ${my_dir}/..api/ ${my_dir}/..api/api.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --go_out=plugins=grpc:${my_dir}/..api"
protoc -I ${my_dir}/../api/ ${my_dir}/../api/api.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --go_out=plugins=grpc:${my_dir}/../api

echo
inf "generating REST gateway stubs..."
inf "protoc -I /usr/local/include/ -I ${my_dir}/../api/ ${my_dir}/../api/api.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${my_dir}/../api"
protoc -I /usr/local/include/ -I ${my_dir}/../api/ ${my_dir}/../api/api.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${my_dir}/../api

echo
inf "generating swagger docs..."
inf "protoc -I /usr/local/include/ -I ${my_dir}/../api/ ${my_dir}/../api/api.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${my_dir}/../swagger"
protoc -I /usr/local/include/ -I ${my_dir}/../api/ ${my_dir}/../api/api.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${my_dir}/../swagger

echo
inf "generating api docs..."
inf "protoc -I api/ -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway api/api.proto --doc_out docs --doc_opt=markdown,api.md"
protoc -I api/ -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway api/api.proto --doc_out docs --doc_opt=markdown,api.md