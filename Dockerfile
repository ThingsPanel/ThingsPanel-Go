# syntax=docker/dockerfile:1
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/app
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build

FROM alpine:latest
ENV TP_PG_IP=172.19.0.4
ENV TP_PG_PORT=5432
ENV TP_MQTT_HOST=172.19.0.5:1883
ENV MQTT_HTTP_HOST=172.19.0.5:8083
ENV TP_REDIS_HOST=172.19.0.6:6379
ENV PLUGIN_HTTP_HOST=172.19.0.8:503
WORKDIR /go/src/app
COPY --from=builder /go/src/app .
EXPOSE 9999
RUN mkdir /docker-entrypoint.d && \
    chmod +x ThingsPanel-Go docker-entrypoint.sh
// 增加预处理过程，方便后期调试，例如独立运行本镜像查看环境等
ENTRYPOINT ["/go/src/app/docker-entrypoint.sh"]
CMD [ "./ThingsPanel-Go" ]
