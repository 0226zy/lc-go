.PHONY: test run build clean

test:
	go test ./... -v

run:
	go run ./cmd/main.go

build:
	go build -o bin/lc-go ./cmd/main.go

clean:
	rm -rf bin/
