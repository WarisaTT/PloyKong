FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server ./cmd/main.go


FROM alpine:3.19
RUN apk --no-cache add ca-certificates wget

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/assets/fonts ./assets/fonts

# ใช้ path นี้กับ volume
RUN mkdir -p /uploads

EXPOSE 8082
HEALTHCHECK --interval=30s --timeout=10s CMD wget -qO- http://localhost:8082/health || exit 1

CMD ["./server"]