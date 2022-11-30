#########
# Proto #
#########
PROTO_SRC_DIR="proto"
PROTO_DEST_DIR="generated"

.PHONY: proto-install
proto-install:
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
