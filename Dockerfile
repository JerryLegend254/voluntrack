FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o main ./cmd/web/*
EXPOSE 8080

ENTRYPOINT ["/app/main"]
