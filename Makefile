run:
	./bin/main

build:
	go build -x -o ./bin/main ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/main

update:
	go build -x -o ./bin/main ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/main

deploy:
	git push git@54.180.147.87:~/blog_hook master
	
	
