.PHONY: clean test
clean:
	rm -rf bin

test:
	go test ./...

build: test
	for CMD in `ls cmd`; do \
		env CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/$$CMD ./cmd/$$CMD/...; \
	done
