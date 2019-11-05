#!/usr/bin/bash

pip install grpcio grpcio-tools

python3 -m grpc_tools.protoc \
-I/project/proto \
-I/project/include \
--python_out=/out \
--grpc_python_out=/out \
/project/proto/*.proto /project/include/scalapb/scalapb.proto

sed -i 's/^\(import.*_pb2\)/from . \1/' /out/*.py
sed -i -e "s/from scalapb/from .scalapb/g" /out/*.py
