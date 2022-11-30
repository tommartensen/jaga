PROTO_SRC_DIR="proto"
PROTO_DEST_DIR="generated"

.PHONY: generate-proto
generate-proto:
	@mkdir -p ${PROTO_DEST_DIR}
	protoc -I=${PROTO_SRC_DIR} \
		--go_out=${PROTO_DEST_DIR} \
		--go-grpc_out=${PROTO_DEST_DIR} \
		${PROTO_SRC_DIR}/**/**/*


TAG=$(shell git describe --tags --abbrev=0)

linker_flags='-s -X main.version=${TAG}'

server:
	go build -ldflags=${linker_flags} -o build/jaga-server ./cmd/jaga-server

cli:
	go build -ldflags=${linker_flags} -o build/jagactl ./cmd/jagactl
