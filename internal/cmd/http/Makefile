NAME=http
VERSION=0.0.1
GOBUILD=../../../_tools/gobuild.py

#----------------------------------------
#      build args
#----------------------------------------
BUILD_ARGS1=--prefix $(NAME)  --version $(VERSION)
BUILD_ARGS2=$(BUILD_ARGS1)


clean:
	python $(GOBUILD) clean $(BUILD_ARGS1)

build-all:clean
	# build windows
	python $(GOBUILD) build $(BUILD_ARGS2) --goos windows --goarch amd64
	# build linux-amd64
	python $(GOBUILD) build $(BUILD_ARGS2) --goos linux --goarch amd64
	# build linux-arm64
	python $(GOBUILD) build $(BUILD_ARGS2) --goos linux --goarch arm64

build-windows:clean
	# build windows
	python $(GOBUILD) build $(BUILD_ARGS2) --goos windows --goarch amd64
