FROM golang:1.13.4-alpine3.10

WORKDIR $GOPATH/src/rest-shell

COPY  ./pkg/ ./pkg/

COPY ./cmd/restshellservice ./cmd/restshellservice

COPY ./vendor/ ./vendor/

COPY ./module/ ./module/

COPY ./docs/ ./docs/

COPY ./test/ ./test/

COPY ./go.mod ./go.mod

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-w -s' -o restshellservice ./cmd/restshellservice

FROM alpine:3.10.3

WORKDIR /restshellservice/app

COPY --from=0 /go/src/rest-shell/restshellservice ./

COPY --from=0 /go/src/rest-shell/docs ./

WORKDIR /restshellservice

ENTRYPOINT ["./app/restshellservice"]

