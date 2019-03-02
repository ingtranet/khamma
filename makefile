DOCKER_APP_NAME=khamma

build :
	docker build -t $(DOCKER_APP_NAME) .

test : build
	docker run --rm $(DOCKER_APP_NAME)