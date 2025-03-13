GO_PROTO_PATH := ./generated

compile_go_protos:
	mkdir -p $(GO_PROTO_PATH)
	protoc \
	--go-grpc_out=$(GO_PROTO_PATH) \
	--go_out=$(GO_PROTO_PATH) \
	-I. ./proto/*.proto
start:
	go run cmd/main.go