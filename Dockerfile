# docker build -t go-app -f Dockerfile .
# docker run  -p 8080:8080 go-app /main

FROM --platform=linux/x86_64 golang:1.18 AS builder

WORKDIR /go/src/go-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM --platform=linux/x86_64 alpine:latest

COPY --from=builder /go/src/go-app/main ./
