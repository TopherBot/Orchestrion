FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o orchestrion ./cmd/orchestrion

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/orchestrion /usr/local/bin/orchestrion
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/orchestrion"]