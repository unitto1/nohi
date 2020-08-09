all: test lint

PKG="gitlab.com/unitto/nohi"
COMMIT="$$(git rev-parse --short HEAD)"

build:
	go build -ldflags "-X ${PKG}/cmd.Commit=${COMMIT}"

test:
	go test -v ./... -cover -race

bench:
	go test -v -benchmem ./... -bench .

full-bench:
	go test -v -cpu 1,2,4,8 -benchmem -benchtime 5s ./... -bench .

dep:
	go get golang.org/x/tools/cmd/goimports
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.30.0

lint:
	golangci-lint run
