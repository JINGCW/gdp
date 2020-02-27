.DEFAULT_GOAL :=all

all: dec-redis dec-mapstructure

dec-deps:
	@mkdir -p bin config

dec-redis: dec-deps
	go build -i -o bin/dec-redis ./src/dego-redis

dec-mapstructure: dec-deps
	go build -i -o bin/dec-mapstructure ./src/demapstructure

build:
	@mkdir -p bin
	go build -o bin/project -v
test:
	go test -v ./...
clean:
	go clean
	@rm -f mybin
run:
	./run.sh
	go build -o mydir -v ./...
	go run mydir/..

make-test:
	mkdir -p bin config && bash --version
