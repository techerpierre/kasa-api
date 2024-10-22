FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/steebchen/prisma-client-go generate
#RUN go build -o api ./cmd/api

EXPOSE 8080

CMD ["go", "run", "./cmd/api"]