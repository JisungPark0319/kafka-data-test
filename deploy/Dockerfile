FROM golang:1.20.5-alpine3.18 as builder
LABEL image.type=builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY . ./

RUN go mod download

RUN go build -o kafka-client .

FROM debian:bullseye-slim
LABEL service.type=nebula

COPY --from=builder /build/kafka-client ./usr/bin

ENTRYPOINT ["kafka-client"]