FROM golang:1.22.1-alpine3.19 AS builder

COPY . /github.com/MirToykin/voice-transcriber-tg-bot/source/
WORKDIR /github.com/MirToykin/voice-transcriber-tg-bot/source/

RUN go mod download
RUN CGO_ENABLED=1 go build -o ./bin/transcriber_bot main.go
RUN chmod +x ./bin/transcriber_bot

FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

WORKDIR /root/
RUN mkdir "storage"
COPY --from=builder /github.com/MirToykin/voice-transcriber-tg-bot/source/bin/transcriber_bot .

CMD ["./transcriber_bot"]