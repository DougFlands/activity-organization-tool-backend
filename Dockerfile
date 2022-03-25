FROM golang:1.17.8-alpine3.15

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/backend
COPY . .

RUN go env && go build -o server .

FROM alpine:latest

WORKDIR /go/src/backend

COPY --from=0 /go/src/backend ./

EXPOSE 8888

ENTRYPOINT ./server -c config.docker.yaml
