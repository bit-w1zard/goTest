.PHONY: build run docker

build:
	go build -o event-service .

run: build
	./event-service

docker:
	docker build -t event-service .
	docker-compose up
