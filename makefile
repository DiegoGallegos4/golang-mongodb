build: 
	docker build --rm -t fitup:latest . -f Dockerfile.prod

build-dev:
	docker build --rm -t fitup.dev:latest . -f Dockerfile.dev

up:
	docker-compose up

down:
	docker-compose down
