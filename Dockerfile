FROM golang:1.19 AS builder

WORKDIR /computer_club

COPY . .

RUN go build -o ./main

FROM alpine:latest

WORKDIR /

COPY --from=builder /computer_club/main main