gen:
	protoc -Iproto --go_out=plugins=grpc:gen proto/*.proto

clean:
	rm pb/*.go

run-server:
	go run cmd/server/main.go -port 8080

run-client:
	go run cmd/client/main.go -address "0.0.0.0:8080"

test:
	go test -cover -race ./...