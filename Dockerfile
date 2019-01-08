FROM golang:latest
WORKDIR /go/src/app
COPY .. .
RUN go build -o web
EXPOSE 6000
CMD ["./web"]
