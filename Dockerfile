FROM golang:latest AS build-env
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod /src/
RUN go mod download
COPY . .
RUN go build -a -o webpalm -trimpath

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
    && rm -rf /var/cache/*

RUN mkdir -p /app \
    && adduser -D webpalm \
    && chown -R webpalm:webpalm /app

USER webpalm
WORKDIR /app

COPY --from=build-env /src/webpalm .

ENTRYPOINT [ "./webpalm" ]