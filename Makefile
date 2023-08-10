all:
	make wire docs fmt build

deps:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest

build:
	go build -o api.exe

linux:
	make wire docs fmt
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -tags lambda -a -ldflags '-extldflags "-static"' -o api.exe

watch:
	air --build.cmd "make" --build.bin "./api.exe"

fmt:
	go fmt ./...
	swag fmt

wire:
	cd api_src && wire

docs:
	swag init

clean:
	rm -rf api.exe

.PHONY: all deps build linux watch fmt wire docs clean
