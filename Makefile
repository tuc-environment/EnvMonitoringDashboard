build:
	go build -o api.exe

watch:
	air --build.cmd "make" --build.bin "./api.exe"

wire:
	wire

clean:
	rm -rf api.exe

.PHONY: build lambda watch wire clean
