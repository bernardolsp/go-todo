FROM golang:1.22

WORKDIR /go/src/app
COPY . .
RUN go get .
RUN go build -o /go/src/app/app
EXPOSE 3000
ENTRYPOINT ["/go/src/app/app"]