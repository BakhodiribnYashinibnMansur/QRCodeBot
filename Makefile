.PHONY:

build:
	go build -o ./.bin/build command/main.go

run: build
	./.bin/build

build-image:
	docker build -t channelhelperbot .

start-container:
	docker run --env-file .env -p 80:80 channelhelperbot