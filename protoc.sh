#!/usr/bin/env bash

set -eu

cd client
PROTOC_GEN_TS_PATH="$(yarn bin)/protoc-gen-ts"
cd ..

PROTO_SRC="./proto"

# server pb生成
# protoc --go_out=plugins=grpc,import_path=proto:. ./server/proto/*.proto

# client pb生成
protoc --ts_out="service=grpc-web:./client/proto" \
    --js_out="import_style=commonjs,binary:./client/proto" \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    -I ${PROTO_SRC} $(find ${PROTO_SRC} -name "*.proto")
