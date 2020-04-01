ROOT_ASSETS ?= ./assets
ROOT_DIST ?= $(ROOT_ASSETS)/login/dist
ROOT_DIST_FILE ?= assets.go

# change out package
ROOT_PACKAG ?= loginComp

all: cleanDist checkTestOSDistPath
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	adm combine js --src=$(ROOT_ASSETS)/src/js/combine/ --dist=$(ROOT_DIST)/all.min.js
	adm combine css --src=$(ROOT_ASSETS)/src/css/ --dist=$(ROOT_DIST)/all.min.css
	adm compile asset --src=$(ROOT_DIST)/ --dist=$$PWD/ --package=$(ROOT_PACKAG)

cleanDist:
	@if [ -f ${ROOT_DIST_FILE} ]; \
	then rm -f ${ROOT_DIST_FILE} && echo "~> removed ${ROOT_DIST_FILE}"; \
	fi
	@if [ -d ${ROOT_DIST} ]; \
	then rm -rf ${ROOT_DIST} && echo "~> removed ${ROOT_DIST}"; \
	fi

checkTestOSDistPath:
	@if [ ! -d ${ROOT_DIST} ]; \
	then mkdir -p ${ROOT_DIST} && echo "~> mkdir ${ROOT_DIST}"; \
	fi
