FROM golang:1.20.6-bullseye AS builder

WORKDIR /app

COPY go.mod /app/
RUN go mod download

COPY . /app
RUN go build -o ./bin/api ./main.go

FROM debian:bullseye-slim

WORKDIR /app
COPY ./entrypoint.sh ./entrypoint.sh
COPY --from=builder /app/bin/api ./bin/api

EXPOSE 8081

ENTRYPOINT [ "./entrypoint.sh" ]
