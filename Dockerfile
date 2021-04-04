FROM golang:alpine AS build
RUN apk --no-cache add make git
WORKDIR /go/src/app
COPY server .
RUN go mod init webserver
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./server

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80