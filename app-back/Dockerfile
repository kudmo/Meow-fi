FROM golang:alpine AS builder

WORKDIR /app-back

ADD go.mod /app-back

COPY .  /app-back

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/main.go

FROM alpine

WORKDIR /app-back

COPY --from=builder /app-back/main /app-back/main


EXPOSE 1323
ENTRYPOINT ["/app-back/main"]