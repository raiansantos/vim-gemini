FROM golang:1.22.3 as builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

COPY server/ .
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]
