FROM golang:1.19
WORKDIR /app/server2
COPY ./server2 .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver2
ENTRYPOINT [ "/app/server2/webserver2" ]