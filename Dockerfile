FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build cmd/api/main.go -o api .

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/api .

EXPOSE 3000

CMD ["./api"]