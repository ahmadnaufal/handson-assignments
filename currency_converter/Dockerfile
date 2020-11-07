FROM alpine:3

RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*

COPY _output/currency-converter /app/currency-converter
COPY .env /app/.env

WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["./currency-converter"]
