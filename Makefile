all:
	make wire docs fmt build

deps:
	go install github.com/google/wire/cmd/wire
	go install github.com/swaggo/swag/cmd/swag

build:
	go build -o api.exe

watch:
	air --build.cmd "make" --build.bin "./api.exe"

fmt:
	go fmt ./...

wire:
	wire

docs:
	swag init

clean:
	rm -rf api.exe

.PHONY: all deps build watch fmt wire docs clean
