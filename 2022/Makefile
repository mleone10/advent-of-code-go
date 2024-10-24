.PHONY: clean test
clean:
	rm -rf newday

test: clean
	go test ./...

build: test
	go build -o newday .
