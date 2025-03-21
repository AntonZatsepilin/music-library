FROM golang:1.23.1-alpine

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ ./

COPY backend/.env .env
COPY backend/schema /app/schema

RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin \
    && go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g cmd/server/main.go -o ./docs

RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]