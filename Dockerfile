FROM iron/go:dev
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .
RUN go build -o web && chmod +x ./web
USER 0
EXPOSE 6000
CMD ./web
