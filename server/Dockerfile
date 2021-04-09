# Stage - Build
FROM golang:alpine AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/todoApp

# Stage - Release
FROM alpine:latest as release
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /go/bin/todoApp /bin/todoApp
EXPOSE 9090
ENTRYPOINT ["todoApp"]