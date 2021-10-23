gen:
	protoc --proto_path=proto  --go_out=. proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go