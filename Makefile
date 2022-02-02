build:
	go build -o ./bin/ ./cmd/...

test:
	bin/mule-server &
	bin/mule-send go.mod &
	bin/mule-recv received_file &
