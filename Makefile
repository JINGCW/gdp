all: test build
build:
	go build -o mybin -v
test:
	go test -v ./...
clean:
	go clean
	rm -f mybin
run:
	./run.sh
	go build -o mydir -v ./...
	go run mydir/..

