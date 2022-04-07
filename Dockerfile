FROM golang:1.18.0-stretch

WORKDIR /app

COPY ./ .

RUN apt update -y && \
    apt-get install -y gcc && \
    go build .

EXPOSE 8080

CMD ["./gotodo"]