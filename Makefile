PROTO_DIR := proto
OUT_DIR := .

.PHONY: all proto user-proto

all: proto

proto: user-proto

user-proto:
	@echo "Generating gRPC code for user-service..."
	protoc -I $(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_DIR)/user/user.proto
