BINARY_NAME=app

.PHONY: all build clean lint docker up up-d

all: dep build

build:
	go build -o ${BINARY_NAME} -v
	chmod +x app

clean:
	rm -f ${BINARY_NAME}

lint:
	golangci-lint run

test:
	go test -v ./...

docker:
	docker-compose build

up:
	docker-compose up

up-d:
	docker-compose up -d

run:
	go run *.go
