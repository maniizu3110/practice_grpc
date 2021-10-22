gen:
	protoc -Iproto --go_out=. proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go