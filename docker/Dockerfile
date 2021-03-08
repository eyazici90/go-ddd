FROM golang:1.14-alpine AS build-env

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . /src

WORKDIR /src

RUN go build -o /src/cmd/http/app /src/cmd/http/main.go

#
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

COPY --from=build-env /src/cmd/http/app /app/main 
COPY --from=build-env /src/cmd/http/config.json /app/config.json 

WORKDIR /app

RUN chmod +x main

ENTRYPOINT ["./main","app"]
