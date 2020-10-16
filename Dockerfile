FROM golang:1.13 as builder

COPY . /go/src/github.com/nadavb/nogobk

WORKDIR /go/src/github.com/nadavb/nogobk/api/server

ENV CGO_ENABLED=0
ENV GO111MODULE=off
ENV GOOD=linux

RUN go build -o /go/bin/nogobk

FROM alpine:3.7

COPY --from=builder /go/bin/nogobk /app/nogobk

WORKDIR /app

ADD api/server/templates templates/
ADD api/server/static static/

CMD /app/nogobk