# this file must use for Postgre

admDBGenerateDemo:
	@echo "Generate ./db/default use config 1-demo.ini"
	cd ./db/default && adm generate -c 1-demo.ini

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
	@echo " isntall adm can use: go install github.com/GoAdminGroup/go-admin/adm"
	@echo "~> make admDBGenerate - adm generate all"
	@echo "~> make admInstall    - adm cli install"
	@echo "~> make admUpdate     - adm cli update"
	@echo ""