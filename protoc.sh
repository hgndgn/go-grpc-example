#!/bin/bash

SRC_DIR="./proto/sorting.proto"
DEST_DIR="Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:."

eval protoc \
  --go_out=$DEST_DIR \
  --go-grpc_out=$DEST_DIR \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  $SRC_DIR