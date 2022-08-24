FROM golang:1.19
WORKDIR /app/server3
COPY ./server3 .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver3
ENTRYPOINT [ "/app/server3/webserver3" ]