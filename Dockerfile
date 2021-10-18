FROM golang:1.16-alpine3.13

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY html/*.html ./

RUN go build -o /online_banking_example

EXPOSE 8080

ENTRYPOINT [ "/online_banking_example" ]