FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
