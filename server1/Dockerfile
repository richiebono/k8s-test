FROM golang:1.19
WORKDIR /app/server1
COPY ./server1 .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver1
ENTRYPOINT [ "/app/server1/webserver1" ]