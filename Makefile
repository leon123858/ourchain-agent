all:
	make start

start:
	go mod tidy
	go fmt ./...
	go run cmd/main.go

build:
	go mod tidy
	go build -o bin/aid cmd/main.go

deploy:
	docker buildx build --no-cache -t go-aid -f ./Dockerfile --platform linux/amd64 .
	docker tag go-aid leon1234858/ourchain-agent
	docker push leon1234858/ourchain-agent

test:
	make unit_test
	make function_test

unit_test:
	go clean -testcache
	go fmt ./...
	go test -v ./...

function_test:
	go fmt ./...
	go run script/contract/runBasicContract.go
	go run script/raw/main.go

doc:
	echo "goto: http://localhost:3000/github.com/leon123858/go-aid"
	pkgsite -http "localhost:3000"

swag:
	swag fmt && cd cmd && swag init

clean:
	go clean