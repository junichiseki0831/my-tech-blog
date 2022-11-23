FROM golang:1.18.4-alpine3.16
RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD ./app /go/src/app

RUN go get -u github.com/go-sql-driver/mysql
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/cosmtrek/air@v1.27.3