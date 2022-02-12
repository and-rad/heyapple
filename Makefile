BINARY_NAME := heyapple
TEST_ARCH := amd64

include .make/nextcloud.mk
include .make/server.mk
include .make/env.mk

build: test-all-server build-server build-nextcloud

run: run-server

test: test-server

install: install-server