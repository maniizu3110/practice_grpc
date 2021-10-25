gen:
	protoc -Iproto --go_out=plugins=grpc:gen proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...