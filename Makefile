.PHONY: check-env prepare clean

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
	if [ -d build ]; then rm -rf build; fi;
