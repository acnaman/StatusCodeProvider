FROM golang:alpine
WORKDIR /opt/www
COPY main.go /opt/www
RUN go build -o app main.go
CMD ["/opt/www/app"]
