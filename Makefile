all:
	make start

start:
	go fmt ./...
	go run cmd/main.go

function_test:
	go fmt ./...
	go run test/function/runBasicContract.go