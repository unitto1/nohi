### Build stage
FROM golang:1.14.7-alpine3.12

ARG flags
WORKDIR /
COPY . .

RUN set -x; CGO_ENABLED=0 go build -ldflags "${flags}" -o /nohi

### Deploy stage
FROM scratch

COPY --from=0 /nohi /usr/bin/nohi

ENTRYPOINT ["nohi"]
CMD ["nohi --help"]
