vendor:
	go mod vendor

test:
	go test ./...

lint:
	golint -set_exit_status ./...

build:
	go build -v -o ideamd .