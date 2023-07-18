FROM golang:latest AS build-env
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod /src/
RUN go mod download
COPY . .
RUN go build -a -o webpalm -trimpath

FROM scratch AS final

COPY --from=build-env /src/webpalm .

ENTRYPOINT [ "./webpalm" ]