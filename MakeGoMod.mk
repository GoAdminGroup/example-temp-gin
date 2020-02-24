# this file must use as base Makefile

modList:
	-GOPROXY="$(ENV_GO_PROXY)" go list -m -json all

modClean:
	@if [ -f ./go.sum ]; then rm -f ./go.sum && echo "~> cleaned file ./go.sum"; else echo "~> has cleaned file ./go.sum"; fi
	@if [ -d ./vendor ]; then rm -rf ./vendor && echo "~> cleaned folder ./vendor"; else echo "~> has cleaned folder ./vendor"; fi

modVerify:
	# in GOPATH must use [ GO111MODULE=on go mod ] to use
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod verify

modDownload:
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod download
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod vendor

modTidy:
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod tidy

dep: modVerify modDownload
	@echo "just check depends info below"

modGraphDependencies:
	GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod graph

modFetch:
	-@go list -m -versions github.com/GoAdminGroup/go-admin | awk '{print $$1 " lastest: " $$NF}'
	-@go list -m -versions github.com/GoAdminGroup/themes | awk '{print $$1 " lastest: " $$NF}'

helpGoMod:
	@echo "Help: MakeGoMod.mk"
	@echo "this project use go mod, so golang version must 1.12+"
	@echo "go mod evn: GOPROXY=$(ENV_GO_PROXY)"
	@echo "~> make dep                  - check depends of project and download all, child task is: modVerify modDownload"
	@echo "~> make modClean             - clearn mod local file and path"
	@echo "~> make modGraphDependencies - see depends graph of this project"
	@echo "~> make modTidy              - tidy depends graph of project"
	@echo "~> make modFetch             - fetch newset version of depends"
	@echo ""