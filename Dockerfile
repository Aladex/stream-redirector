FROM golang:1.17
COPY . /app
WORKDIR /app
RUN go build -ldflags "-w -s" -o bin/stream-redirector cmd/main.go
EXPOSE 8080
CMD ./bin/stream-redirector
