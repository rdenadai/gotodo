FROM golang:1.18.1-alpine3.15

WORKDIR /app

COPY ./ .

RUN apk update && \
    apk add build-base git && \
    go build .

EXPOSE 8080

CMD ["./gotodo"]