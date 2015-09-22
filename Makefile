test:
	go test -race -v -cpu 1,2 ./...

.PHONY: test
