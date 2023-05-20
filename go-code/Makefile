gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=pb --go_out=pb

clean:
	find . -name "*.go" -type f -delete

run:
	go run main.go

test:
	go test -cover -race ./...

