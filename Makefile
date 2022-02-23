.PHONY: build
build:
	go build -o ./build/pong .

run:
	./build/pong