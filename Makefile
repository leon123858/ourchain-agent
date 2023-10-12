all:
	make start

start:
	go fmt ./...
	go run cmd/main.go

build:
	go mod tidy
	go build -o bin/aid cmd/main.go

function_test:
	go fmt ./...
	go run test/function/runBasicContract.go

doc:
	echo "goto: http://localhost:3000/github.com/leon123858/go-aid"
	pkgsite -http "localhost:3000"

uml:
	echo "should install `go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest` first"
	echo "preview in: https://www.plantuml.com/"
	echo "tool git url: https://github.com/jfeliu007/goplantuml"
	goplantuml -recursive ./ > UML.puml

clean:
	go clean