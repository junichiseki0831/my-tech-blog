# https://hub.docker.com/_/golang
FROM golang:1.18.1-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY ./app ./
RUN go mod tidy
RUN go mod download
RUN go build main.go

CMD ["./main"]