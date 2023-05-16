FROM golang:latest

WORKDIR /usr/src/app

# RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

RUN go build -o /bin/main ./cmd && ls

# CMD [ "./bin/main" ]
FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=build /bin/main /usr/src/app/