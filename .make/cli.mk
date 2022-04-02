build-cli:
	@rm -rf ./out/cli
	@CGO_ENABLED=0 GOARCH=amd64 go build -o ./out/cli/${BINARY_NAME}-cli-amd64 github.com/and-rad/heyapple/cmd/cli
