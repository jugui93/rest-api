FROM golang:latest AS build

WORKDIR /usr/src/app

# RUN go install github.com/cosmtrek/air@latest

COPY cmd/. ./cmd/
COPY database/. ./database/
COPY handlers/. ./handlers/
COPY models/. ./models/
COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/main ./cmd

# CMD [ "./bin/main" ]
FROM alpine:latest

WORKDIR /bin

COPY --from=build /bin/main /bin