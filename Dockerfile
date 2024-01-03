FROM golang:latest AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY .. .
ENV CGO_ENABLED=0 GO111MODULE=on
RUN go build -o ./bin ./main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/bin/* ./
ENTRYPOINT ["./main"]