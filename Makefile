include .env

PROJECT_PATH=$(shell pwd)
MODULE_NAME=cr-api

BUILD_NUM_FILE=./deploy/build_num.txt
BUILD_NUM=$$(cat ./deploy/build_num.txt)
APP_VERSION=$$(cat ./deploy/version.txt)
TARGET_VERSION=$(APP_VERSION).$(BUILD_NUM)

TARGET_DIR=bin
OUTPUT=$(PROJECT_PATH)/$(TARGET_DIR)/$(MODULE_NAME)
MAIN_DIR=/main.go
LDFLAGS=-X main.BUILD_TIME=`date -u '+%Y-%m-%d_%H:%M:%S'`
LDFLAGS+=-X main.GIT_HASH=`git rev-parse HEAD`
LDFLAGS+=-s -w

all: config docker-build docker-push

config:
	@if [ ! -d $(TARGET_DIR) ]; then mkdir $(TARGET_DIR); fi

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(OUTPUT) $(PROJECT_PATH)$(MAIN_DIR)

build-copy:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(OUTPUT) $(PROJECT_PATH)$(MAIN_DIR)
	cp $(OUTPUT) ./

docker-build:
	@echo "TARGET_VERSION : $(IMAGE_REPOSITORY):$(TARGET_VERSION)"
	docker build -f Dockerfile --tag $(IMAGE_REPOSITORY):$(TARGET_VERSION) .

docker-push:
	@echo "TARGET_VERSION : $(IMAGE_REPOSITORY):$(TARGET_VERSION)"
	docker push $(IMAGE_REPOSITORY):$(TARGET_VERSION)

docker-release:
	@echo "TARGET_VERSION : $(IMAGE_REPOSITORY):$(TARGET_VERSION)"
	docker build -f Dockerfile --tag $(IMAGE_REPOSITORY):latest .
	docker push $(IMAGE_REPOSITORY):latest

docker-run:
	@echo "TARGET_VERSION : $(IMAGE_REPOSITORY):$(TARGET_VERSION)"
	docker run \
	--name cr-api \
	-p 8088:8088 \
	-it --rm $(IMAGE_REPOSITORY):$(TARGET_VERSION)

ecr-access:
	bash -c ./deploy/ecr/ecr_access.sh

target-version:
	@echo "========================================"
	@echo "APP_VERSION    : $(APP_VERSION)"
	@echo "BUILD_NUM      : $(BUILD_NUM)"
	@echo "TARGET_VERSION : $(IMAGE_REPOSITORY):$(TARGET_VERSION)"
	@echo "========================================"

build-num:
	@echo $$(($$(cat $(BUILD_NUM_FILE)) + 1 )) > $(BUILD_NUM_FILE)

swagger-go-path:
	export PATH=$(go env GOPATH)/bin:$PATH

swagger-set:
	swag init

clean:
	rm -f $(PROJECT_PATH)/$(TARGET_DIR)/$(MODULE_NAME)*