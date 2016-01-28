test:
	go test -race -v -cpu 1,2,4 ./...

bench:
	go test -cpu 1,2,4 -bench . ./...

.PHONY: test bench
