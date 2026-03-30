FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rest-api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/rest-api .
EXPOSE 8080
CMD ["./rest-api"]
