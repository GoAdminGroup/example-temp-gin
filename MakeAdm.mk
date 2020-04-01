# this file must use for Postgre

ENV_ROOT_ADM_GEN_PATH ?= ./docker/depend/adm
ENV_ROOT_ADM_GEN_DEFAULT ?= $(ENV_ROOT_ADM_GEN_PATH)/default

admDBGenerateDemo:
	@if [ ! -d "./tables/demo" ]; \
	then mkdir -p "./tables/demo" && echo "~> mkdir ./tables/demo"; \
	fi
	@echo "Generate $(ENV_ROOT_ADM_GEN_DEFAULT) use config 1-demo.ini"
	cd $(ENV_ROOT_ADM_GEN_DEFAULT) && adm generate -c 1-demo.ini

admDBGenerateDefault:
	@if [ ! -d "./tables/config" ]; \
	then mkdir -p "./tables/config" && echo "~> mkdir ./tables/config"; \
	fi
	@echo "Generate $(ENV_ROOT_ADM_GEN_DEFAULT) use config 1-config.ini"
	cd $(ENV_ROOT_ADM_GEN_DEFAULT) && adm generate -c 1-config.ini

admInstall:
	-go install -v github.com/GoAdminGroup/go-admin/adm
	adm -V

admUpdate:
	-go get -v -u github.com/GoAdminGroup/go-admin/adm
	adm -V

admDBGenerate: admDBGenerateDefault
	@echo "this can gen more database"

helpAdm:
	@echo "Help: MakeAdm.mk"
	@echo "this project use adm can"
	@echo "install adm can use: go install github.com/GoAdminGroup/go-admin/adm"
	@echo "~> make admDBGenerate - adm generate all"
	@echo "~> make admInstall    - adm cli install"
	@echo "~> make admUpdate     - adm cli update"
	@echo ""
