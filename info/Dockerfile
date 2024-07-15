FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN apt-get update

RUN go mod download
RUN go build -o info-golang ./backend/cmd/main.go

EXPOSE 8081

CMD ["./info-golang"]
