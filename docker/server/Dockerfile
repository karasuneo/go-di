FROM golang:1.22.0-alpine

WORKDIR /app

COPY /src /app
COPY /.env /app/.env

CMD ["go", "run", "main.go"]
