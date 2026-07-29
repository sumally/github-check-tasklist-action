FROM golang:1.26 AS builder

WORKDIR /go/src/app
COPY . .
RUN go build -o action ./

FROM gcr.io/distroless/base-debian13

COPY --from=builder /go/src/app/action /action
ENTRYPOINT ["/action"]
