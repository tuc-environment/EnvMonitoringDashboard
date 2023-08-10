all:
	make wire docs fmt build

deps:
	cd _webapp && npm i
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest

build:
	cd _webapp && npm run build
	go build -o EnvMonitoringDashboard.exe

linux:
	make wire docs fmt
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -tags lambda -a -ldflags '-extldflags "-static"' -o EnvMonitoringDashboard.exe

watch:
	air --build.cmd "make" --build.bin "./EnvMonitoringDashboard.exe"

fmt:
	cd _webapp && npm run format
	go fmt ./...
	swag fmt

wire:
	cd api_src && wire

docs:
	swag init -o api_src/docs

clean:
	rm -rf EnvMonitoringDashboard.exe

.PHONY: all deps build linux watch fmt wire docs clean
