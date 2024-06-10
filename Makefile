build:
	go build -o ./build/bin/proxy_server

run: build
	./build/bin/proxy_server
