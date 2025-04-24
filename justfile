build:
    mkdir -p build
    go build -o build/fbmirros main.go

clean:
    rm -rf build

check:
    go vet ./...
    go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint run ./...

fmt:
    go fmt ./...

test:
    go test -v ./...

