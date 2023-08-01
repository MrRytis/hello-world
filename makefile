run:
	TZ=Europe/Vilnius \
	APP_NAME="hello-world" \
	ENV="develop" \
  	APP_VERSION="0.0.1" \
	go run main.go

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build --no-cache

ssh:
	docker-compose exec app sh