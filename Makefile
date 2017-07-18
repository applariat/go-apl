
ALL_OS := linux darwin windows
ALL_ARCH := amd64

VERSION ?= unknown
ARCH ?= amd64
OS ?= linux

all: all-build

build-%:
	@$(MAKE) --no-print-directory OS=$* build

all-build: $(addprefix build-, $(ALL_OS))

build: bin/apl-$(VERSION)-$(OS)_$(ARCH)

bin/apl-$(VERSION)-$(OS)_$(ARCH): build-dirs
	@echo "building: $@"
	@VERSION=$(VERSION) OS=$(OS) ARCH=$(ARCH) ./build/build.sh

check:
	@./build/check.sh cmd pkg

clean:
	rm -rf bin

build-dirs:
	@mkdir -p bin
