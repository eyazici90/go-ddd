FROM golang:1.17-alpine AS build-env

ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o /build/cmd/http/app /build/cmd/http/main.go

#
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

COPY --from=build-env /build/cmd/http/app /app/main 
COPY --from=build-env /build/cmd/http/config.json /app/config.json 

WORKDIR /app

RUN chmod +x main

ENTRYPOINT ["./main","app"]
