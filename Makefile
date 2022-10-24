build:
	go run main.go

run: build
	./app/messaging

watch:
	reflex -s -r '\.go$$' make run

proto:
	protoc -I ./protobuff/ ./protobuff/*.proto --go_out=./rpcs --go-grpc_out=./rpcs