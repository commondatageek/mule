build:
	go build -o ./bin/ ./cmd/...

test:
	bin/mule-server &
	sleep 2
	bin/mule-send -host localhost -port 8881 -infile go.mod &
	sleep 2
	bin/mule-recv -host localhost -port 8882 -outfile received_file &
