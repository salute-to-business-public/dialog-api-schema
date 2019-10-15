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

touch /out/__init__.py
touch /out/scalapb/__init__.py

sed -i 's/^\(import.*_pb2\)/from . \1/' /out/*.py
sed -i -e "s/from scalapb/from .scalapb/g" /out/*.py
