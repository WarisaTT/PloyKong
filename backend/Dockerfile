FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git ca-certificates

WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/server ./cmd/main.go


FROM alpine:3.19
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/assets/fonts ./assets/fonts

# ใช้ path นี้กับ volume
RUN mkdir -p /uploads

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=10s CMD wget -qO- http://localhost:3000/health || exit 1

CMD ["./server"]