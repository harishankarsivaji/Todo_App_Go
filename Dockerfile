FROM golang:alpine AS builder
COPY server /build
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o . .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build .

EXPOSE 9090
ENTRYPOINT /app --port 9090