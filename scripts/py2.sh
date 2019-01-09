#!/usr/bin/bash

pip install grpcio grpcio-tools

python3 -m grpc_tools.protoc \
-I/project/proto \
-I/project/include \
-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--python_out=/out \
--grpc_python_out=/out \
/project/proto/*.proto

python3 -m grpc_tools.protoc \
-I/project/include \
-I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--python_out=/out \
--grpc_python_out=/out \
/project/include/scalapb/scalapb.proto

touch /out/scalapb/__init__.py
