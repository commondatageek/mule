build: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd64/ ./cmd/...
	cd cmd/mule && go install

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o ./bin/linux/arm64/ ./cmd/...

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/amd64/ ./cmd/...

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin/arm64/ ./cmd/...
