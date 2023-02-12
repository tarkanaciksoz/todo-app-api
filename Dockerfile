FROM golang:1.19-alpine
ARG ENV
ARG BIND_ADDRESS
WORKDIR /app

RUN apk add -q --update \
    && apk add -q \
            bash \
            git \
            curl \
            nano \
    && rm -rf /var/cache/apk/*

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN APP_ENV=$ENV go build -o main main.go
RUN CGO_ENABLED=0 APP_ENV=$ENV go test -v ./...

EXPOSE $BIND_ADDRESS
CMD APP_ENV=$ENV ./main