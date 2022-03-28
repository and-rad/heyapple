BINARY_NAME := heyapple
TEST_ARCH := amd64

include .make/web.mk
include .make/server.mk
include .make/env.mk

build: test-server build-web build-server

run: run-server

test: test-all-server

install: install-server