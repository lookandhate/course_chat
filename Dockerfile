FROM golang:1.22-alpine as builder
ENV CGO_ENABLED=0

WORKDIR /app

COPY . .
RUN go mod download


WORKDIR /app
RUN go build -o /bin/chat.go ./cmd/chat.go/main.go

FROM alpine

WORKDIR /bin

COPY --from=builder /bin/chat /bin/chat
COPY .env .env
COPY ./config ./config

CMD ["/bin/chat"]