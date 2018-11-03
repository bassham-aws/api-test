FROM golang:1.11 as builder

COPY . /go/src/github.com/bassham-aws/api-test
WORKDIR /go/src/github.com/bassham-aws/api-test

EXPOSE 8089

RUN CGO_ENABLED=0 go build -v -o api-test

FROM alpine:latest

COPY --from=builder /go/src/github.com/bassham-aws/api-test/api-test /app/

CMD /app/api-test
