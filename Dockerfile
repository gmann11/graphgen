# build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# test
FROM alpine:latest as test
WORKDIR /app
COPY --from=builder /app/main .
COPY words .
ENTRYPOINT [ "/app/main" ]

# deploy
FROM alpine:latest as deploy
WORKDIR /app
COPY --from=builder /app/main .
COPY words .
ENTRYPOINT [ "/app/main" ]
