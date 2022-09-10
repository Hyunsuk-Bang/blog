
generate:
	protoc -Iproto --go_opt=module=blog --go_out=. \
	--go-grpc_opt=module=blog --go-grpc_out=. proto/*.proto

build:
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

PHONY: generate, build