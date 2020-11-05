FROM golang:1.13

WORKDIR $GOPATH/src/github.com/reaperhero/go-gin-example
COPY . $GOPATH/src/github.com/reaperhero/go-gin-example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]