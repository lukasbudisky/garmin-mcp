NAME="garmin-mcp"

build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -a -o "${NAME}"
