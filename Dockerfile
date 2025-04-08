FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

CMD ["go", "run", "main.go"]