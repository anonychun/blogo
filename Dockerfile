FROM golang:latest AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -ldflags="-s -w" -o _output/bin/server main.go

FROM ubuntu:latest
RUN apt update
RUN apt install -y ca-certificates
WORKDIR /app
COPY --from=builder /go/src/app/_output/bin/server .
COPY --from=builder /go/src/app/docs .
COPY --from=builder /go/src/app/migrations ./migrations

ENTRYPOINT ["./server"]
CMD ["launch"]