FROM golang:1.17

WORKDIR /go/src/app
COPY . .
RUN go build -o action ./
ENTRYPOINT ["/go/src/app/action"]
