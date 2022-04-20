build-cli:
	@rm -rf ./out/cli
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/cli/heyapple-cli-linux-amd64 github.com/and-rad/heyapple/cmd/cli
