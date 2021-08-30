build:
	go build -o notification app/api/main.go

run: build
	./notification

build-grpc-gateway:
	go build grpc_gateway.go

grpc-gateway: build-grpc-gateway
	./grpc_gateway