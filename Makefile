NAME="garmin-mcp"

build-container:
	container build --platform=linux/amd64 -f Dockerfile .

build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -a -o "${NAME}-darwin-arm64"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o "${NAME}-linux-amd64"
