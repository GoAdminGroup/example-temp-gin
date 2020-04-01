# this file must use as base Makefile

SERVER_TEST_SSH_ALIASE = aliyun-ecs
SERVER_TEST_FOLDER = /home/work/Document/
SERVER_REPO_SSH_ALIASE = temp-goadmin-base
SERVER_REPO_FOLDER = /home/ubuntu/$(ROOT_NAME)

checkTestBuildPath:
	@if [ ! -d ${ROOT_TEST_BUILD_PATH} ]; \
	then mkdir -p ${ROOT_TEST_BUILD_PATH} && echo "~> mkdir ${ROOT_TEST_BUILD_PATH}"; \
	fi

checkTestDistPath:
	@if [ ! -d ${ROOT_TEST_DIST_PATH} ]; \
	then mkdir -p ${ROOT_TEST_DIST_PATH} && echo "~> mkdir ${ROOT_TEST_DIST_PATH}"; \
	fi

checkTestOSDistPath:
	@if [ ! -d ${ROOT_TEST_OS_DIST_PATH} ]; \
	then mkdir -p ${ROOT_TEST_OS_DIST_PATH} && echo "~> mkdir ${ROOT_TEST_OS_DIST_PATH}"; \
	fi

checkReleaseDistPath:
	@if [ ! -d ${ROOT_REPO_DIST_PATH} ]; \
	then mkdir -p ${ROOT_REPO_DIST_PATH} && echo "~> mkdir ${ROOT_REPO_DIST_PATH}"; \
	fi

checkReleaseOSDistPath:
	@if [ ! -d ${ROOT_REPO_OS_DIST_PATH} ]; \
	then mkdir -p ${ROOT_REPO_OS_DIST_PATH} && echo "~> mkdir ${ROOT_REPO_OS_DIST_PATH}"; \
	fi

distTest: dep buildMain checkTestBuildPath
	mv ./build/main $(ROOT_TEST_BUILD_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_BUILD_PATH)
	-cp -R ./static $(ROOT_TEST_BUILD_PATH)
	@echo "=> pkg at: $(ROOT_TEST_BUILD_PATH)"

distTestOS: dep buildARCH checkTestOSDistPath
	@echo "=> Test at: $(ENV_DIST_OS) ARCH as: $(ENV_DIST_ARCH)"
	mv ./build/main $(ROOT_TEST_OS_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_OS_DIST_PATH)
	-cp -R ./static $(ROOT_TEST_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_OS_DIST_PATH)"

distTestOSTar: distTestOS
	@echo "=> start tar test as os $(ENV_DIST_OS) $(ENV_DIST_ARCH)"
	tar zcvf $(ROOT_DIST)/$(ENV_DIST_OS)/test/$(ROOT_NAME)-$(ENV_DIST_OS)-$(ENV_DIST_ARCH)-$(ENV_DIST_VERSION).tar.gz $(ROOT_TEST_OS_DIST_PATH)

distRelease: dep buildMain checkReleaseDistPath
	mv ./build/main $(ROOT_REPO_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_DIST_PATH)
	-cp -R ./static $(ROOT_REPO_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_DIST_PATH)"

distReleaseOS: dep buildARCH checkReleaseOSDistPath
	@echo "=> Release at: $(ENV_DIST_OS) ARCH as: $(ENV_DIST_ARCH)"
	mv ./build/main $(ROOT_REPO_OS_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_OS_DIST_PATH)
	-cp -R ./static $(ROOT_REPO_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_OS_DIST_PATH)"

distReleaseOSTar: distReleaseOS
	@echo "=> start tar release as os $(ENV_DIST_OS) $(ENV_DIST_ARCH)"
	tar zcvf $(ROOT_DIST)/$(ENV_DIST_OS)/release/$(ROOT_NAME)-$(ENV_DIST_OS)-$(ENV_DIST_ARCH)-$(ENV_DIST_VERSION).tar.gz $(ROOT_REPO_OS_DIST_PATH)

scpDistReleaseOSTar: distReleaseOSTar
	scp $(ROOT_DIST)/$(ENV_DIST_OS)/release/$(ROOT_NAME)-$(ENV_DIST_OS)-$(ENV_DIST_ARCH)-$(ENV_DIST_VERSION).tar.gz $(SERVER_REPO_SSH_ALIASE):$(SERVER_REPO_FOLDER)
	@echo "=> must check below config of set for relase OS Scp"

scpDockerComposeTest:
	scp ./conf/test/docker-compose.yml $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)
	@echo "=> finish update docker compose at test"

helpDist:
	@echo "Help: helpDist.mk"
	@echo "-- distTestOS or distReleaseOS will out abi as: $(ENV_DIST_OS) $(ENV_DIST_ARCH) --"
	@echo "~> make distTest         - build dist at $(ROOT_TEST_DIST_PATH) in local OS"
	@echo "~> make distTestOS       - build dist at $(ROOT_TEST_OS_DIST_PATH) as: $(ENV_DIST_OS) $(ENV_DIST_ARCH)"
	@echo "~> make distTestOSTar    - build dist at $(ROOT_TEST_OS_DIST_PATH) as: $(ENV_DIST_OS) $(ENV_DIST_ARCH) and tar"
	@echo "~> make distRelease      - build dist at $(ROOT_REPO_DIST_PATH) in local OS"
	@echo "~> make distReleaseOS    - build dist at $(ROOT_REPO_OS_DIST_PATH) as: $(ENV_DIST_OS) $(ENV_DIST_ARCH)"
	@echo "~> make distReleaseOSTar - build dist at $(ROOT_REPO_OS_DIST_PATH) as: $(ENV_DIST_OS) $(ENV_DIST_ARCH) and tar"
	@echo ""
