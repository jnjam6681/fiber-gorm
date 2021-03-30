.PHONY: clean build

APP_NAME = goapp
BUILD_DIR = $(PWD)/build

newman:
	newman run ./api/Fiber-Gorm.postman_collection.json \
	-e ./api/Fiber-Gorm.postman_environment.json

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

clean:
	rm -rf ./build

docker.up:
	docker-compose up -d --build

docker.down:
	docker-compose down