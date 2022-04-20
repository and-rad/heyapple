DEV_ARCH := amd64
DEV_OS := linux

include .make/web.mk
include .make/cli.mk
include .make/server.mk
include .make/env.mk

build: test-server build-web build-server build-cli

run: run-server

test: test-all-server
