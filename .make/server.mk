NUM_TESTS := `find -name "*_test.go"|xargs grep "{ //"|wc -l`
TEST_PACKAGES := `go list ./internal/... | grep -v -e /defaults/ -e /mock -e /web`

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
	@CGO_ENABLED=0 GOOS=${DEV_OS} GOARCH=${DEV_ARCH} go build -o ./out/server/heyapple-${DEV_OS}-${DEV_ARCH} github.com/and-rad/heyapple/cmd/server

run-server:
	@./out/server/heyapple-${DEV_OS}-${DEV_ARCH}

test-server:
	@echo "testing ${NUM_TESTS} cases:"
	@go test -short -cover -p 1 -timeout 30m ${TEST_PACKAGES}

test-all-server:
	@echo "testing ${NUM_TESTS} cases:"
	@go test -cover -race -p 1 -timeout 30m ${TEST_PACKAGES}
