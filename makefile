all: test build
build:
    ## linux
	env GOOS=linux GOARCH=amd64 go build -o sample-code-server app/main.go
	
	## mac as development
	go build -o dev app/main.go
test: 
	go test -v ./...
