NAME := tweet-gateway
TAG := latest

all:

build: build-protoc
	DOCKER_BUILDKIT=1 docker build -t ${NAME}:${TAG} .

build-protoc:
	python -m grpc_tools.protoc --proto_path=. --python_out=. --grpc_python_out=. tweet-gateway.proto

run:
	docker run -it --rm -p 31337:31337 -v $(shell pwd):/tmp -e TG_CONFIG=/tmp/.secret ${NAME}:${TAG}
