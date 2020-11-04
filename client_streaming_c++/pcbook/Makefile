gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*go

server:
	go run cmd/server/main.go

client:
	go run cmd/client/libVNetClient.go

test:
	go test -cover -race ./...