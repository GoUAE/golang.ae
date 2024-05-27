ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/golang.ae
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /golang.ae .


FROM debian:bookworm

WORKDIR /usr/src/golang.ae
COPY --from=builder /golang.ae /usr/local/bin/
COPY --from=builder /usr/src/golang.ae/public ./public/
CMD ["golang.ae"]
