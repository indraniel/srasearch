.PHONY: check-env prepare clean

# srasearch version
DEBVERSION := 0.1.0

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

# only meant to be invoked on a Linux/Debian based system
debian: srasearch
	test -d build || mkdir build
	echo 2.0 > build/debian-binary
	echo "Package: srasearch" > build/control
	echo "Version: 1.0-${DEBVERSION}" >> build/control
	echo "Architecture: amd64" >> build/control
	echo "Section: science" >> build/control
	echo "Maintainer: indraniel <indraniel@gmail.com>" >> build/control
	echo "Priority: optional" >> build/control
	echo "Description: An NCBI Short Read Archive Upload Management search utility" >> build/control
	mkdir -p build/usr/local/bin
	cp srasearch build/usr/local/bin
#	tar cvzf build/data.tar.gz --owner=0 --group=0 -C build usr
	tar cvzf build/data.tar.gz -C build usr
	tar cvzf build/control.tar.gz -C build control
	cd build && ar rc srasearch-${DEBVERSION}.deb \
		debian-binary control.tar.gz data.tar.gz \
		&& cd ..
