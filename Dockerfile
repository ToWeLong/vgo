FROM golang:1.16.9 AS builder
WORKDIR /go/src
COPY . .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.io,direct \
    && CGO_ENABLED=0 go build -o App github.com/towelong/vgo/cmd/app

FROM alpine AS final
MAINTAINER welong <towelong@qq.com>
WORKDIR /app
COPY --from=builder /go/src/App /app/App
RUN chmod a+xr -R /app/App
COPY --from=builder /go/src/configs /app/configs
EXPOSE 8081
ENTRYPOINT ["/app/App", "-conf", "/app/configs/config.prod.yaml"]
