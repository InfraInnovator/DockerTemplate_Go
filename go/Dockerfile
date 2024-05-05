FROM golang:latest

COPY src/app.go /app.go
RUN /usr/local/go/bin/go build /app.go
ENTRYPOINT ["/go/app"]
