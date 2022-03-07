FROM golang:1.17 as builder

WORKDIR /go/src/app
COPY . .
RUN go build -o action ./

FROM gcr.io/distroless/base-debian11

COPY --from=builder /go/src/app/action /action
ENTRYPOINT ["/action"]
