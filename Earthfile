# Earthfile

FROM golang:alpine
WORKDIR /code

ARG APP=aws-cdk_demo

init:
    RUN go mod init github.com/MatthiasScholz/$APP
    SAVE ARTIFACT go.mod AS LOCAL go.mod

shell:
    COPY ./ ./
    RUN --interactive-keep sh

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod

build:
    FROM +deps
    COPY main.go .
    RUN go build -o build/$APP main.go
    SAVE ARTIFACT build/$APP /$APP AS LOCAL build/$APP
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

run:
    FROM +build
    COPY build/$APP .
    COPY data/test.md .
    RUN ./$APP convert test.md

docker:
    COPY +build/$APP .
    ENTRYPOINT ["/code/$APP"]
    SAVE IMAGE $APP:latest
