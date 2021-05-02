.PHONY: all

PROJ_PATH := ${CURDIR}
DOCKER_PATH := ${PROJ_PATH}/docker

APP=bot

BASIC_IMAGE=default
IMAGE_POSTFIX=-image

build-app:
	go build -o .bin/${APP} cmd/${APP}/main.go
	chmod ugo+x .bin/${APP}

build-docker:
	docker build -t ${BASIC_IMAGE} -f ${DOCKER_PATH}/builder.Dockerfile .
	docker build -t ${APP}${IMAGE_POSTFIX} -f ${DOCKER_PATH}/${APP}.Dockerfile .

app-setup-and-up: build-docker app-up

app-up: build-app
	docker-compose up

all: app-setup-and-up

app-bot-bash:
	docker-compose run --rm --no-deps --name bot_service ${APP} bash
