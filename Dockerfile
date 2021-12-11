FROM golang:1.17.5-alpine3.15 as builder

MAINTAINER dingjiefeng

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release

WORKDIR /go/src/web

COPY . /go/src/web

RUN go build -o app main.go

FROM alpine:3.15 as prod

MAINTAINER dingjiefeng

ENV GIN_MODE=debug \
    PORT=8080

WORKDIR /root/

EXPOSE 8080:$PORT

COPY --from=0 /go/src/web/app .

CMD ["./app"]
