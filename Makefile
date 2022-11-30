#########
# Proto #
#########
PROTOC_VERSION=3.20.3
PROTO_SRC_DIR="proto"
PROTO_DEST_DIR="generated"

# Support different OS to diff local and CI
ifeq ($(shell uname -s),Linux)
PROTOC_URL = https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip
endif
ifeq ($(shell uname -s),Darwin)
PROTOC_URL = https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-osx-x86_64.zip
endif

.PHONY: proto-install
proto-install:
	@echo "Install protoc $(PROTOC_VERSION)"
	@mkdir -p $(GOPATH)/bin
	@curl $(PROTOC_URL) -sLo /tmp/protoc.zip
	@unzip -o -d /tmp /tmp/protoc.zip bin/protoc
	install /tmp/bin/protoc $(GOPATH)/bin/protoc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: proto-generate
proto-generate:
	@mkdir -p ${PROTO_DEST_DIR}
	protoc -I=${PROTO_SRC_DIR} \
		--go_out=${PROTO_DEST_DIR} \
		--go-grpc_out=${PROTO_DEST_DIR} \
		${PROTO_SRC_DIR}/**/**/*

#########
# Build #
#########
TAG=$(shell git describe --tags --abbrev=0)
linker_flags='-s -X main.version=${TAG}'

.PHONY: build-cli
build-cli:
	go build -ldflags=${linker_flags} -o build/jagactl ./cmd/jagactl

.PHONY: build-server
build-server:
	go build -ldflags=${linker_flags} -o build/jaga-server ./cmd/jaga-server
