run:
	go run main.go

build:
	go build main.go

clean:
	go mod tidy 
	go fmt ./...