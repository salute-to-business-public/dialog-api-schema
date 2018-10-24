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
echo "import six
PATH_WORKAROUND = six.PY3

if PATH_WORKAROUND:
    import sys
    import os
    sys.path.append(os.path.dirname(__file__))
" > /out/__init__.py
touch /out/scalapb/__init__.py
