FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o fcc-request-header-parser-ms

EXPOSE 3000

CMD ./fcc-request-header-parser-ms