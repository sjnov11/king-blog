run:
	./bin/main

build:
	go build -x -o ./bin/main ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/main

update:
	echo "Update to server..."
	go build -x -o ./bin/main ./src
	sudo setcap CAP_NET_BIND_SERVICE=+eip ./bin/main

deploy:
	echo "Deploy to server..."
	git push git@54.180.105.111:~/blog_hook master
	
	
