FROM golang:1.25.3-trixie

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
