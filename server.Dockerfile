# Server Dockerfile
FROM golang:1.22.1

RUN mkdir /fileguard
WORKDIR /fileguard

COPY . ./
RUN go mod download

RUN go build -o /go-docker-fileguard ./server

EXPOSE 8080

CMD ["/go-docker-fileguard"]
