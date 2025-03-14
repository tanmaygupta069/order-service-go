GO_PROTO_PATH := ./generated

compile_go_protos:
	mkdir -p $(GO_PROTO_PATH)
	protoc \
		--go_out=$(GO_PROTO_PATH) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_PROTO_PATH) --go-grpc_opt=paths=source_relative \
		-I proto \
		proto/common/common.proto \
		proto/order/order.proto \
		proto/holding/holding.proto
start:
	go run cmd/main.go