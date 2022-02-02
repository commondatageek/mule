build:
	go build -o ./bin/ ./cmd/...

test:
	bin/mule-server &
	sleep 2
	bin/mule-send go.mod &
	sleep 2
	bin/mule-recv received_file &
