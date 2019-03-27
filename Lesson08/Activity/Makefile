all: build docker-build docker-push

build:
	GOOS=linux GARCH=amd64 go build -o app main.go

docker-build:
	docker build -t suakbas/lesson08:v1 .

docker-push:
	docker push suakbas/lesson08:v1