FROM golang:1.19-alpine
ARG ENV
ARG BIND_ADDRESS
WORKDIR /app

RUN apk add -q --update \
        bash \
        git \
        curl \
        nano \
        && rm -rf /var/cache/apk/*

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN APP_ENV=$ENV go build -o main main.go

EXPOSE $BIND_ADDRESS
CMD APP_ENV=$ENV ./main