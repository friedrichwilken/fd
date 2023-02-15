run:
	go run main.go

build:
	go build main.go

test:
	go test ./...

clean:
	go mod tidy 
	go fmt ./...
