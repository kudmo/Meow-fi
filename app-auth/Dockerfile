FROM golang:alpine AS builder

WORKDIR /app-auth

ADD go.mod /app-auth

COPY .  /app-auth

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/main.go

FROM alpine

WORKDIR /app-auth

COPY --from=builder /app-auth/main /app-auth/main


EXPOSE 1323
ENTRYPOINT ["/app-auth/main"]