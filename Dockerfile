FROM golang:1.22.3-alpine AS builder

WORKDIR /jcli

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

RUN make build

EXPOSE 8080

CMD ["./build/jcli"]

