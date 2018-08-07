.PHONY: docs java js swagger gateway python

PROJECT_PATH   := $(shell pwd)
PROTO_PATH     := $(PROJECT_PATH)/proto
JAVA_GEN_PATH  := $(PROJECT_PATH)/java
JS_GEN_PATH    := $(PROJECT_PATH)/js
DOCS_PATH      := $(PROJECT_PATH)/docs
SWAGGER_PATH   := $(PROJECT_PATH)/swagger
HTTP_GW_PATH   := $(PROJECT_PATH)/http_gateway
PYTHON_PATH    := $(PROJECT_PATH)/python

python:
	mkdir -p $(PYTHON_PATH)
	docker run --name py_grpc_builder -d -i -t \
	-v $(PROJECT_PATH):/project \
	-v $(GOPATH):/go \
	-v $(PYTHON_PATH):/out \
	-w /project \
	python:latest bash
	docker exec -it -w /project py_grpc_builder bash scripts/py.sh
	docker rm -f py_grpc_builder

java:
	mkdir -p $(JAVA_GEN_PATH)
	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--plugin=protoc-gen-grpc-java=opt/protoc-gen-grpc-java-1.13.2-osx-x86_64.exe \
  --grpc-java_out="$(JAVA_GEN_PATH)" \
	--java_out="$(JAVA_GEN_PATH)" \
	$(PROTO_PATH)/*.proto

	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	--java_out="$(JAVA_GEN_PATH)" \
	include/scalapb/scalapb.proto

js:
	mkdir -p $(JS_GEN_PATH)
	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--js_out=$(JS_GEN_PATH) \
	$(PROTO_PATH)/*.proto

swagger:
	mkdir -p $(SWAGGER_PATH)
	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	-I/usr/local/include \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:$(SWAGGER_PATH) \
	$(PROTO_PATH)/*.proto

gateway:
	mkdir -p $(HTTP_GW_PATH)
	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:$(HTTP_GW_PATH) \
	$(PROTO_PATH)/*.proto
	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,\
	Mscalapb/scalapb.proto=github.com/gogo/protobuf/types,\
	plugins=grpc:$(HTTP_GW_PATH) \
	$(PROTO_PATH)/*.proto

docs:
	rm -rf $(DOCS_PATH)
	mkdir -p $(DOCS_PATH)
	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	--plugin=protoc-gen-doc=$(GOPATH)/bin/protoc-gen-doc \
  --doc_out=$(DOCS_PATH) \
	--doc_opt=markdown,readme.md \
	$(PROTO_PATH)/*.proto

	@protoc -I$(PROTO_PATH) \
	-Iinclude \
	--plugin=protoc-gen-doc=$(GOPATH)/bin/protoc-gen-doc \
  --doc_out=$(DOCS_PATH) \
	--doc_opt=html,index.html \
	$(PROTO_PATH)/*.proto
