# this file must load as base Makefile

ROOT_DOCKER_DEPEND_MK_FOLDER ?= docker/depend# do not change
ROOT_DOCKER_COMPOSE_DEPEND_DEV ?= docker/$(ROOT_NAME)-dev# do not change

dockerDependDevFileInit:
	@if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then echo "-> now open ENV_NEED_PROXY"; \
	fi
	cd $(ROOT_DOCKER_DEPEND_MK_FOLDER) && \
	if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then bash dev-services.sh -p -b $(ENV_DIST_VERSION) -b $(ENV_DIST_VERSION) -n $(ROOT_NAME); \
	else bash dev-services.sh -b $(ENV_DIST_VERSION) -n $(ROOT_NAME); \
	fi

dockerDependDevUp:
	cd $(ROOT_DOCKER_COMPOSE_DEPEND_DEV) && \
	docker-compose up -d --remove-orphans
	@echo "=> if contain permission error can use [ make dockerDependDevFixPermission ] to fix"
	@echo "=> init db can use: [ docker exec -it $(ROOT_NAME)-mysql bash ]"
	@echo "and to path: [ cd /tmp/db ] "
	@echo "init admin-table and data: [ bash init-admin.sh ]"
	@echo "init biz table and data: [ bash maintain-biz.sh ]"

dockerDependDevStop:
	cd $(ROOT_DOCKER_COMPOSE_DEPEND_DEV) && \
	docker-compose stop

dockerDependDevContainRemove:
	cd $(ROOT_DOCKER_COMPOSE_DEPEND_DEV) && \
	docker-compose rm -s -f -v

dockerDependDevRestart: dockerDependDevStop dockerDependDevFileInit dockerDependDevUp
	@echo "=> restart finish, please check contain run status"

dockerDependDevFixPermission:
	@echo "=> fix permission: $(ROOT_DOCKER_COMPOSE_DEPEND_DEV)/data/redis-$(ROOT_NAME)"
	sudo chown -R 1001:1001 $(ROOT_DOCKER_COMPOSE_DEPEND_DEV)/data/redis-$(ROOT_NAME)

helpDockerDepend:
	@echo "Help: helpDockerDepend.mk"
	@echo "~> make dockerDependDevFileInit      - init Docker config at folder $(ROOT_DOCKER_DEPEND_MK_FOLDER) and init dev"
	@echo "~> make dockerDependDevUp            - update depend at folder $(ROOT_DOCKER_COMPOSE_DEPEND_DEV)"
	@echo "~> make dockerDependDevStop          - stop depend at folder $(ROOT_DOCKER_COMPOSE_DEPEND_DEV)"
	@echo "~> make dockerDependDevContainRemove - stop depend at folder $(ROOT_DOCKER_COMPOSE_DEPEND_DEV)"
	@echo "~> make dockerDependDevRestart       - stop depend at folder $(ROOT_DOCKER_COMPOSE_DEPEND_DEV) and dockerDependDevFileInit then dockerDependDevUp"
	@echo ""
