NUM_TESTS := `find -name "*_test.go"|xargs grep "{ //"|wc -l`
INSTALL_DIR := /tmp/heyapple
CONFIG_DIR := /tmp/heyapple

build-web:
	@npm run build --prefix ./web/src/login

build-server:
	@rm -rf ./out/server
	@CGO_ENABLED=0 GOARCH=amd64 go build -o ./out/server/${BINARY_NAME}-amd64 heyapple/cmd/web

run-server:
	@./out/server/${BINARY_NAME}-${TEST_ARCH}

test-server:
	@echo "testing ${NUM_TESTS} cases:"
	@go test -short -cover -p 1 -timeout 30m ./pkg/...

test-all-server:
	@echo "testing ${NUM_TESTS} cases:"
	@go test -cover -p 1 -timeout 30m ./pkg/...

install-server:
	@cp ./out/server/${BINARY_NAME}-amd64 ${INSTALL_DIR}/
	@cp ./configs/sample.env ${CONFIG_DIR}/heyapple.env
	@if [ -e .env ] ;then cp .env ${CONFIG_DIR}/heyapple.env ;fi
