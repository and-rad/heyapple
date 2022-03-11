BINARY_NAME := heyapple
TEST_ARCH := amd64

include .make/nextcloud.mk
include .make/server.mk
include .make/env.mk

build: test-server build-web build-server build-nextcloud

run: run-server

test: test-all-server

install: install-server