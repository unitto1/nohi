### Build stage
FROM golang:1.14.7-alpine3.12

ARG flags
WORKDIR /
COPY . .

RUN set -x; apk add --no-cache upx \
    && CGO_ENABLED=0 go build -ldflags "-s ${flags}" -o nohi \
    && upx -9 nohi

### Deploy stage
FROM scratch

COPY --from=0 /nohi /usr/bin/nohi

ENTRYPOINT ["nohi"]
CMD ["nohi --help"]
