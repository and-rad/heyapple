NUM_TESTS := `find -name "*_test.go"|xargs grep "{ //"|wc -l`
INSTALL_DIR := /tmp/heyapple
CONFIG_DIR := /tmp/heyapple

build-server:
	@rm -rf ./out/server
	@if command -v scour &> /dev/null; then \
		scour ./assets/icons.svg ./internal/web/static/img/icons.svg \
			--enable-id-stripping \
			--protect-ids-noninkscape \
			--remove-descriptive-elements \
			--enable-comment-stripping \
			--strip-xml-prolog \
			--strip-xml-space; fi
	@CGO_ENABLED=0 GOARCH=amd64 go build -o ./out/server/${BINARY_NAME}-amd64 github.com/and-rad/heyapple/cmd/web

run-server:
	@./out/server/${BINARY_NAME}-${TEST_ARCH}

test-server:
	@echo "testing ${NUM_TESTS} cases:"
	@go test -short -cover -p 1 -timeout 30m ./internal/...

test-all-server:
	@echo "testing ${NUM_TESTS} cases:"
	@go test -cover -race -p 1 -timeout 30m ./internal/...

install-server:
	@cp ./out/server/${BINARY_NAME}-amd64 ${INSTALL_DIR}/
	@cp ./configs/sample.env ${CONFIG_DIR}/heyapple.env
	@if [ -e .env ] ;then cp .env ${CONFIG_DIR}/heyapple.env ;fi
