build-cli:
	@rm -rf ./out/cli
	@CGO_ENABLED=0 GOOS=${DEV_OS} GOARCH=${DEV_ARCH} go build -o ./out/cli/heyapple-cli-${DEV_OS}-${DEV_ARCH} github.com/and-rad/heyapple/cmd/cli
