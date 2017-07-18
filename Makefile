.PHONY: build

all: build

SRC_DIRS := cmd pkg
PKG := github.com/applariat/go-apl
VERSION ?= unknown

ALL_OS := linux darwin windows

check:
	@./build/check.sh $(SRC_DIRS)

build: $(addprefix build-, $(ALL_OS))

build-%: build-dirs
	PKG=$(PKG) VERSION=$(VERSION) OS=$* ARCH=amd64 ./build/build.sh

clean:
	rm -rf bin

build-dirs:
	@mkdir -p bin