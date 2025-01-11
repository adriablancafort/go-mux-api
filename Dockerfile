FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/api .

EXPOSE 8000

CMD ["./api"]