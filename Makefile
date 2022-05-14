.PHONY: build
build:
	go build -o build/pts -v ./...

.PHONY:test
test:
	go test -v ./...

.PHONY: docker-build
docker-build: build
	docker build -t pts:latest -f Dockerfile .
