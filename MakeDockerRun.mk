# this file must use as base Makefile

# use Makefile ROOT_NAME
ROOT_DOCKER_CONTAINER_NAME ?= $(ROOT_NAME)
ROOT_DOCKER_CONTAINER_PORT ?= 39000
# change this for docker parent
ROOT_DOCKER_IMAGE_PARENT_NAME ?= golang
ROOT_DOCKER_IMAGE_PARENT_TAG ?= 1.13.3-stretch
# change this for dockerRunLinux or dockerRunDarwin
ROOT_DOCKER_IMAGE_NAME ?= $(ROOT_NAME)
# can change as local set or read Makefile DIST_VERSION
ROOT_DOCKER_IMAGE_TAG ?= $(DIST_VERSION)

# For Docker dev images init task
initDockerDevImages:
	@echo "~> start init this project in docker"
	@echo "-> check version"
	go version
	@echo "-> check env golang"
	go env
	@echo "-> install swag"
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod download
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod vendor

dockerLocalImageInit:
	docker build --tag $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG) .

dockerLocalImageRebuild:
	-docker image rm $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	docker build --tag $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG) .

dockerDevBuild:
	docker image inspect --format='{{ .Created}}' $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	ENV_WEB_PORT=$(ROOT_DOCKER_CONTAINER_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_IMAGE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_MAKE_TASK=buildMain \
	docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_CONTAINER_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_CONTAINER_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_CONTAINER_NAME) -f"

localIPLinux:
	@echo "=> now run as docker with linux"
	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_LINUX)"

dockerRunLinux: localIPLinux
	docker image inspect --format='{{ .Created}}' $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_LINUX) \
	ENV_WEB_PORT=$(ROOT_DOCKER_CONTAINER_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_IMAGE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_MAKE_TASK=dev \
	docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_CONTAINER_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_CONTAINER_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_CONTAINER_NAME) -f"

dockerRestartLinux: localIPLinux
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_IMAGE_NAME)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_LINUX) \
	ENV_WEB_PORT=$(ROOT_DOCKER_CONTAINER_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_IMAGE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_MAKE_TASK=dev \
	docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_CONTAINER_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_CONTAINER_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_CONTAINER_NAME) -f"

localIPDarwin:
	@echo "=> now run as docker with darwin"
	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_DARWIN)"

dockerRunDarwin: localIPDarwin
	docker image inspect --format='{{ .Created}}' $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_DARWIN) \
	ENV_WEB_PORT=$(ROOT_DOCKER_CONTAINER_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_IMAGE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_MAKE_TASK=dev \
	docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_CONTAINER_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_CONTAINER_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_CONTAINER_NAME) -f"

dockerRestartDarwin: localIPDarwin
	docker inspect --format='{{ .State.Status}}' $(ROOT_NAME)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_DARWIN) \
	ENV_WEB_PORT=$(ROOT_DOCKER_CONTAINER_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_IMAGE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_MAKE_TASK=dev \
	docker-compose restart
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_CONTAINER_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_CONTAINER_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_CONTAINER_NAME) -f"

dockerStop:
	ROOT_NAME=$(ROOT_DOCKER_IMAGE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_WEB_PORT=$(ROOT_DOCKER_CONTAINER_PORT) \
	docker-compose stop

dockerPrune: dockerStop
	ROOT_NAME=$(ROOT_DOCKER_IMAGE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_WEB_PORT=$(ROOT_DOCKER_CONTAINER_PORT) \
	docker-compose rm -f $(ROOT_DOCKER_IMAGE_NAME)
	-docker rmi -f $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	docker network prune
	docker volume prune

helpDockerRun:
	@echo "Help: MakeDockerRun.mk"
	@echo "Before run this project in docker must use"
	@echo "~> make dockerLocalImageInit to init Docker image"
	@echo "or use"
	@echo "~> make dockerLocalImageRebuild to rebuild Docker image"
	@echo "After build Docker image success"
	@echo "~> make dockerRunLinux  - run docker-compose server as $(ROOT_DOCKER_IMAGE_NAME):$(DIST_VERSION) \
	container-name at $(ROOT_DOCKER_CONTAINER_NAME) in dockerRunLinux"
	@echo "~> make dockerRunDarwin - run docker-compose server as $(ROOT_DOCKER_IMAGE_NAME):$(DIST_VERSION) \
	container-name at $(ROOT_DOCKER_CONTAINER_NAME) in macOS"
	@echo "~> make dockerStop      - stop docker-compose container-name at $(ROOT_DOCKER_CONTAINER_NAME)"
	@echo "~> make dockerPrune     - stop docker-compose container-name at $(ROOT_DOCKER_CONTAINER_NAME) and try to remove"
	@echo ""