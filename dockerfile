FROM golang:1.20-alpine AS builder
WORKDIR /opentelemetry&tracing

COPY go.mod ./
COPY go.sum ./

ENV NEW_RELIC_API_KEY=5e8afd086390d507ce709c6c92da1e9bf8f0NRAL

RUN go mod download
RUN go mod tidy

COPY . .
RUN apk add --no-cache git
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

