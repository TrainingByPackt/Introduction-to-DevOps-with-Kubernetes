all: build docker-build docker-push

build:
	GOOS=linux GARCH=amd64 go build -o v1/v1 v1/v1.go
	GOOS=linux GARCH=amd64 go build -o v2/v2 v2/v2.go

docker-build:
	docker build -t suakbas/lesson07:v1 v1/
	docker build -t suakbas/lesson07:v2 v2/

docker-push:
	docker push suakbas/lesson07:v1
	docker push suakbas/lesson07:v2