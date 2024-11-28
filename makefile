dev:
	go build && ./gonflict

build:
	go build -ldflags="-s -w"
