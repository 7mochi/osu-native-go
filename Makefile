ifndef PLATFORM
ifeq ($(OS),Windows_NT)
    PLATFORM := win-x64
    LIB_NAME := osu.Native.dll
else
    UNAME_S := $(shell uname -s)
    UNAME_M := $(shell uname -m)
    ifeq ($(UNAME_S),Linux)
        ifeq ($(UNAME_M),aarch64)
            PLATFORM := linux-arm64
        else
            PLATFORM := linux-x64
        endif
        LIB_NAME := osu.Native.so
    endif
    ifeq ($(UNAME_S),Darwin)
        PLATFORM := osx-arm64
        LIB_NAME := osu.Native.dylib
    endif
endif
endif

BUILD_DIR := build
NATIVE_DIR := native/bin

.PHONY: all build-osu-native copy-native test test-short clean

all: build-osu-native copy-native

build-osu-native:
	dotnet publish osu-native/osu.Native -c Release -r $(PLATFORM) -o $(BUILD_DIR)/generated
	python3 scripts/fix_cabinet_header.py $(BUILD_DIR)/generated/cabinet.h
	cp $(BUILD_DIR)/generated/cabinet.h $(NATIVE_DIR)/

copy-native:
	mkdir -p $(NATIVE_DIR)/$(PLATFORM)
ifeq ($(LIB_NAME),osu.Native.dll)
	cp $(BUILD_DIR)/generated/$(LIB_NAME) $(NATIVE_DIR)/$(PLATFORM)/
else
	cp $(BUILD_DIR)/generated/$(LIB_NAME) $(NATIVE_DIR)/$(PLATFORM)/lib$(LIB_NAME)
endif

test:
	go test -v ./...

test-short:
	go test -short -v ./...

clean:
	rm -rf $(BUILD_DIR) $(NATIVE_DIR)
