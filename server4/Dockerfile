FROM golang:1.19
WORKDIR /app/server4
COPY ./server4 .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver4
ENTRYPOINT [ "/app/server4/webserver4" ]