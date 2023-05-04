FROM golang:1.18.0-alpine3.14

WORKDIR /opentelemetry&tracing

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .
RUN apk add --no-cache git
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

