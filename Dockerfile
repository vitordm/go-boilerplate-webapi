FROM golang:1.25-alpine AS builder

ENV APPLICATION_PORT=":8080"

WORKDIR /app

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o application

EXPOSE 8080

FROM alpine:latest

EXPOSE 80
EXPOSE 8080

COPY --from=builder /app/application /usr/bin/application

ENTRYPOINT [ "/usr/bin/application" ]


