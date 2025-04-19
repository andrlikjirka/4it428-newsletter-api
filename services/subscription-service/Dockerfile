# -------- Stage 1: Dev & Build --------
FROM golang:1.24-alpine as dev

# Install Air for dev hot reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.work ./
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o subscription-service ./services/subscription-service/cmd/api

EXPOSE 8082

# -------- Stage 2: Production --------
FROM alpine:latest AS prod

WORKDIR /app
COPY --from=dev /app/subscription-service .

CMD ["./subscription-service"]
