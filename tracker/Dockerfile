FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o tracker-golang ./backend/cmd/main.go

RUN chmod +x scripts/wait-for-postgres.sh

CMD ["./tracker-golang"]