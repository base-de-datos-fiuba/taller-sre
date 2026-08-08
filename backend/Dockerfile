FROM golang:1.24-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -o /bytemarket ./cmd/server

FROM alpine:3.22
RUN adduser -D -H app
USER app
COPY --from=build /bytemarket /bytemarket
EXPOSE 8080
ENTRYPOINT ["/bytemarket"]
