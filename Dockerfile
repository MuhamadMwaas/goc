# Builder stage for the main application
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/application cmd/api/main.go

# Builder stage for the migration tool
FROM golang:1.25-alpine AS builder-migrate

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/migrate cmd/migrate/main.go

# Final stage for the main application
FROM alpine:latest AS app

WORKDIR /app

COPY --from=builder /app/application .
COPY migration ./migration

EXPOSE 8080

CMD ["./application"]

# Final stage for running migrations
FROM alpine:latest AS migrate

WORKDIR /app

COPY --from=builder-migrate /app/migrate .
COPY migration ./migration

CMD ["./migrate"]