NUM_TESTS := `find -name "*_test.go"|xargs grep "{ //"|wc -l`

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
