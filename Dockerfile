FROM golang:latest

WORKDIR /go
COPY src ./src

ENV USER root
ENV DATABASE test
ENV HOST mariadb

EXPOSE 8080

RUN go build src/main.go

ENTRYPOINT ["./main"]