FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN apt-get update

RUN go mod download
RUN go build -o tracker-app ./backend/cmd/main.go

CMD ["./tracker-app"]