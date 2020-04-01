.PHONY: test check clean build dist all

TOP_DIR := $(shell pwd)

# ifeq ($(FILE), $(wildcard $(FILE)))
# 	@ echo target file not found
# endif

# each tag change this
ENV_DIST_VERSION := v1.0.0
# need open proxy 1 is need 0 is default
ENV_NEED_PROXY=1

ENV_DIST_OS := linux
ENV_DIST_ARCH := amd64

ENV_DIST_OS_DOCKER ?= linux
ENV_DIST_ARCH_DOCKER ?= amd64

ROOT_NAME ?= example-temp-gin

# ignore used not matching mode
ROOT_TEST_INVERT_MATCH ?= "vendor|handler\/biz"
# set ignore of test case like grep -v -E "vendor|fataloom" to ignore vendor and fataloom package
ROOT_TEST_LIST ?= $$(go list ./... | grep -v -E $(ROOT_TEST_INVERT_MATCH))
# test max time
ROOT_TEST_MAX_TIME := 1m

ROOT_BUILD_PATH ?= ./build
ROOT_LOG_PATH ?= ./log
ROOT_ROOT_DATA_PATH ?= ./data
ROOT_SWAGGER_PATH ?= ./docs
ROOT_DIST ?= $$PWD/dist
ROOT_REPO ?= $$PWD/dist

ROOT_TEST_BUILD_PATH ?= $(ROOT_BUILD_PATH)/test/$(ENV_DIST_VERSION)
ROOT_TEST_DIST_PATH ?= $(ROOT_DIST)/test/$(ENV_DIST_VERSION)
ROOT_TEST_OS_DIST_PATH ?= $(ROOT_DIST)/$(ENV_DIST_OS)/test/$(ENV_DIST_VERSION)
ROOT_REPO_DIST_PATH ?= $(ROOT_REPO)/$(ENV_DIST_VERSION)
ROOT_REPO_OS_DIST_PATH ?= $(ROOT_REPO)/$(ENV_DIST_OS)/release/$(ENV_DIST_VERSION)

ROOT_LOCAL_IP_V4_LINUX = $$(ifconfig enp8s0 | grep inet | grep -v inet6 | cut -d ':' -f2 | cut -d ' ' -f1)
ROOT_LOCAL_IP_V4_DARWIN = $$(ifconfig en0 | grep inet | grep -v inet6 | cut -d ' ' -f2)

# can use as https://goproxy.cn/
# https://goproxy.io/
# https://gocenter.io
# https://mirrors.aliyun.com/goproxy/
ENV_GO_PROXY ?= https://goproxy.cn/

# include MakeDockerRun.mk for docker run
include MakeGoMod.mk
include MakeDockerDepend.mk
include MakeDockerRun.mk
include MakeAdm.mk
include MakeDist.mk

checkEnvGOPATH:
ifndef GOPATH
	@echo Environment variable GOPATH is not set
	exit 1
endif

init:
	@echo "~> start init this project"
	@echo "-> check go version"
	go version
	@echo "-> check go env"
	go env
	@if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then echo "-> now open ENV_NEED_PROXY then use GOPROXY=$(ENV_GO_PROXY)"; \
	fi
	-@if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod vendor; \
	else GO111MODULE=on go mod vendor; \
	fi
	@echo "-> init finish"
	@echo "~> you can use [ make help ] see more task"

cleanBuild:
	@if [ -d ${ROOT_BUILD_PATH} ]; \
	then rm -rf ${ROOT_BUILD_PATH} && echo "~> cleaned ${ROOT_BUILD_PATH}"; \
	else echo "~> has cleaned ${ROOT_BUILD_PATH}"; \
	fi

cleanDist:
	@if [ -d ${ROOT_DIST} ]; \
	then rm -rf ${ROOT_DIST} && echo "~> cleaned ${ROOT_DIST}"; \
	else echo "~> has cleaned ${ROOT_DIST}"; \
	fi

cleanLog:
	@if [ -d ${ROOT_LOG_PATH} ]; \
	then rm -rf ${ROOT_LOG_PATH} && echo "~> cleaned ${ROOT_LOG_PATH}"; \
	else echo "~> has cleaned ${ROOT_LOG_PATH}"; \
	fi

cleanRootData:
	@if [ -d ${ROOT_ROOT_DATA_PATH} ]; \
	then rm -rf ${ROOT_ROOT_DATA_PATH} && echo "~> cleaned ${ROOT_ROOT_DATA_PATH}"; \
	else echo "~> has cleaned ${ROOT_ROOT_DATA_PATH}"; \
	fi

clean: cleanBuild cleanLog cleanRootData
	@echo "~> clean finish"

checkLogPath:
	@if [ ! -d ${ROOT_LOG_PATH} ]; then mkdir -p ${ROOT_LOG_PATH} && echo "~> mkdir ${ROOT_LOG_PATH}"; fi

buildMain: dep
	@echo "-> start build local OS"
	@if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then echo "-> now use GOPROXY=$(ENV_GO_PROXY)"; \
	fi
	@if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then GOPROXY="$(ENV_GO_PROXY)" go build -o build/main main.go; \
	else go build -o build/main main.go; \
	fi

buildARCH: dep
	@echo "-> start build OS:$(ENV_DIST_OS) ARCH:$(ENV_DIST_ARCH)"
	@if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then echo "-> now use GOPROXY=$(ENV_GO_PROXY)"; \
	fi
	@if [ $(ENV_NEED_PROXY) -eq 1 ]; \
	then GOPROXY="$(ENV_GO_PROXY)" GOOS=$(ENV_DIST_OS) GOARCH=$(ENV_DIST_ARCH) go build -tags netgo -o build/main main.go; \
	else GOOS=$(ENV_DIST_OS) GOARCH=$(ENV_DIST_ARCH) go build -tags netgo -o build/main main.go; \
	fi

dev: buildMain
	-ENV_WEB_AUTO_HOST=true ./build/main -c ./conf/config.yaml

test:
	@echo "=> run test start"
	#=> go test -test.v $(ROOT_TEST_LIST)
	@go test -test.v $(ROOT_TEST_LIST)

testFail:
	@echo "=> run test start"
	#=> go test -test.v $(ROOT_TEST_LIST)
	@go test -test.v $(ROOT_TEST_LIST) | grep FAIL

testBenchmem:
	@echo "=> run test benchmem start"
	@go test -test.benchmem

cloc:
	# https://stackoverflow.com/questions/26152014/cloc-ignore-exclude-list-file-clocignore
	cloc --exclude-list-file=.clocignore .

helpProjectRoot:
	@echo "Help: Project root Makefile"
	@echo ""
	@echo "-- now build name: $(ROOT_NAME) version: $(ENV_DIST_VERSION) --"
	@echo "~> make init         - check base env of this project"
	@echo "~> make clean        - remove binary file and log files"
	@echo "~> make test         - run test case all"
	@echo "~> make testBenchmem - run go test benchmem case all"
	@echo "~> make dev          - run server use conf/config.yaml"

help: helpGoMod helpDockerDepend helpDockerRun helpAdm helpProjectRoot
	@echo ""
	@echo "-- more info see Makefile include: MakeGoMod.mk MakeDockerRun.mk --"
