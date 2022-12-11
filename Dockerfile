# https://hub.docker.com/_/golang
FROM golang:1.18.1-alpine

EXPOSE 8080

RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./app /go/src/app

RUN go get -u github.com/go-sql-driver/mysql && \
    go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go install github.com/cosmtrek/air@v1.27.3

CMD ["air"]