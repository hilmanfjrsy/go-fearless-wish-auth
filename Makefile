run:
	go run main.go
docker-up:
	docker-compose up -d
docker-down:
	docker-compose down
docker-start:
	docker-compose start go-fearless-wish-auth
docker-stop:
	docker-compose stop