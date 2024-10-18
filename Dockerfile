FROM golang:1.20-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/steebchen/prisma-client-go

RUN go build -o api ./cmd/api/main.go

FROM alpine:3.18

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/api .

EXPOSE 8080

CMD ["./api"]