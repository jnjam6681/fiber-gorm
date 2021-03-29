.PHONY: clean test security build run

newman:
	newman run ./api/Fiber-Gorm.postman_collection.json \
	-e ./api/Fiber-Gorm.postman_environment.json

docker.up:
	docker-compose up -d --build

docker.down:
	docker-compose down