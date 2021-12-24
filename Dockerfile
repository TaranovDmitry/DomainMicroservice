FROM golang:1.17

WORKDIR /app

ADD . /app

WORKDIR /app

RUN go mod download
RUN go build -o main .

CMD ["/app/main"]
