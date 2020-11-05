FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8282

RUN go build

CMD ["./master"]    