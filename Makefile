all: test lint

PKG="gitlab.com/unitto/nohi"
COMMIT="$$(git rev-parse --short HEAD)"

dep:
	python3 -m pip install -r build/requirements.dev.txt
	go get golang.org/x/tools/cmd/goimports
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.30.0

### Go shortcuts
build: test
	CGO_ENABLED=0 go build -ldflags "-X ${PKG}/cmd.Commit=${COMMIT}"

test:
	CGO_ENABLED=0 go test -v ./... -cover -race

bench:
	CGO_ENABLED=0 go test -v -benchmem ./... -bench .

bench-long:
	CGO_ENABLED=0 go test -v -cpu 1,2,4,8 -benchmem -benchtime 5s ./... -bench .

lint:
	golangci-lint -v run


### Docs shortcuts
docs:
	mkdocs build --verbose --site-dir public

docs-serve:
	mkdocs serve


### Docker shortcuts
docker:
	docker build --build-arg flags="-X ${PKG}/cmd.Commit=${COMMIT}" -t unitto/nohi:dev -f build/Dockerfile .

docker-test:
	docker run unitto/nohi:dev
