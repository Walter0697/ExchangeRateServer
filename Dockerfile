FROM golang:1.19 AS build-stage

WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
ENV CGO_ENABLED=0

COPY . .
RUN go get
RUN go build -o /app/Server

FROM alpine:3.9 AS production-stage
WORKDIR /app
COPY --from=build-stage /app/Server /app/Server

EXPOSE 80

CMD ["/app/Server"]