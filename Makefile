run:
	./bin/main

build:
	go build -x -o ./bin/main ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/main
