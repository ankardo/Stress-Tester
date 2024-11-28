FROM golang:1.23.3 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod tidy && \
  go build -o stress-tester ./cmd/cli/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/stress-tester /app/stress-tester

ENTRYPOINT ["/app/stress-tester"]
