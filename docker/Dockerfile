FROM golang:1.22.1-alpine AS build-env

ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o /build/cmd/go-ddd/app /build/cmd/go-ddd/main.go

#
FROM gcr.io/distroless/static-debian12

COPY --from=build-env /build/cmd/go-ddd/app /app/main
COPY --from=build-env /build/cmd/go-ddd/config.json /app/config.json

WORKDIR /app

ENTRYPOINT ["./main","app"]
