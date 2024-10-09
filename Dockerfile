### Step 1: Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-http-server


### Step 2: Runtime stage
FROM scratch

COPY --from=builder /app/go-http-server /

EXPOSE 8080

ENTRYPOINT ["/go-http-server"]
