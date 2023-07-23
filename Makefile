all:
	make wire docs fmt build

deps:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest

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
	swag fmt

clean:
	rm -rf api.exe

.PHONY: all deps build watch fmt wire docs clean
