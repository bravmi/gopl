
test:
	go test ./... | grep -v 'no test'

lint:
	golangci-lint run
