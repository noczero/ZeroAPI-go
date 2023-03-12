FROM golang:1.20-alpine

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

RUN mkdir /app
COPY . /app
WORKDIR /app

# for live reloading
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o main cmd/main.go" -command="/app/main"
