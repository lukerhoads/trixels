FROM golang:1.17-alpine
RUN apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
COPY ./server ./server
COPY ./abigen ./abigen
RUN go build -o /trixels-server
EXPOSE 8080
CMD ["/trixels-server"] 