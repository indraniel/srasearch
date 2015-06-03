.PHONY: check-env prepare clean

# debian
DEB_BUILD_DIR := "deb-build"
DEB_RELEASE_VERSION := "1"

SOURCES=$(wilcard  *.go **/*.go **/**/*.go)
GODEP := $(GOPATH)/bin/godep

all: srasearch

dev: export PATH := $(GOPATH)/bin/:$(PATH) 
dev: check-env $(SOURCES)
	go-bindata -debug -o assets/assets.go \
		-pkg="assets" \
		-ignore=web/static/js/bootstrap.js \
		-ignore=web/static/js/jquery.js \
		web/static/... web/views/...
	go build -o srasearch-dev main.go

srasearch: export PATH := $(GOPATH)/bin/:$(PATH) 
srasearch: check-env $(SOURCES)
	go-bindata -o assets/assets.go \
		-pkg="assets" \
		-ignore=web/static/js/bootstrap.js \
		-ignore=web/static/js/jquery.js \
		web/static/... web/views/...
	go build -o srasearch main.go

prepare: check-env
	@echo "GOPATH is: ${GOPATH}"
	@echo "GOROOT is: ${GOROOT}"
	go get github.com/jteeuwen/go-bindata/go-bindata
	go get github.com/tools/godep
	$(GODEP) restore

check-env:
ifndef GOROOT
	$(error environment variable GOROOT is undefined)
endif

ifndef GOPATH
	$(error environment variable GOPATH is undefined)
endif

clean:
	if [ -e srasearch-dev ]; then rm srasearch-dev; fi;
	if [ -e srasearch ]; then rm srasearch; fi;
	if [ -e assets/assets.go ]; then rm -rf assets; fi;
	if [ -d $(DEB_BUILD_DIR) ]; then rm -rf $(DEB_BUILD_DIR); fi;

# only meant to be invoked on a Linux/Debian based system
debian: srasearch
	$(eval DEB_PKG_VERSION := $(shell ./srasearch version))
	test -d $(DEB_BUILD_DIR) || mkdir $(DEB_BUILD_DIR)
	echo 2.0 > $(DEB_BUILD_DIR)/debian-binary
	echo "Package: srasearch" > $(DEB_BUILD_DIR)/control
	echo "Version: ${DEB_PKG_VERSION}-${DEB_RELEASE_VERSION}" >> $(DEB_BUILD_DIR)/control
	echo "Architecture: amd64" >> $(DEB_BUILD_DIR)/control
	echo "Section: science" >> $(DEB_BUILD_DIR)/control
	echo "Maintainer: indraniel <indraniel@gmail.com>" >> $(DEB_BUILD_DIR)/control
	echo "Priority: optional" >> $(DEB_BUILD_DIR)/control
	echo "Description: An NCBI Short Read Archive Upload Management search utility" >> $(DEB_BUILD_DIR)/control
	mkdir -p $(DEB_BUILD_DIR)/usr/local/bin
	cp srasearch $(DEB_BUILD_DIR)/usr/local/bin
	tar cvzf $(DEB_BUILD_DIR)/data.tar.gz --owner=0 --group=0 -C $(DEB_BUILD_DIR) usr
	tar cvzf $(DEB_BUILD_DIR)/control.tar.gz -C $(DEB_BUILD_DIR) control
	cd $(DEB_BUILD_DIR) && ar rc srasearch_${DEB_PKG_VERSION}-${DEB_RELEASE_VERSION}.deb \
		debian-binary control.tar.gz data.tar.gz \
		&& cd ..
