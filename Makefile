SERVICE_NAME=Example
SERVICE_TAG=latest

build-app:
	docker build . -f Dockerfile-local -t ${SERVICE_NAME}:${SERVICE_TAG}

test:
	go test ./...
