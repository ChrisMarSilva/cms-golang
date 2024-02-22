## Build
FROM golang:1.22.0 AS builder
#FROM golang:1.22.0-alpine AS builder 
#FROM golang:1.22.0-alpine3.19 AS builder 
WORKDIR /app
USER ${USER}
RUN apt-get autoremove && apt-get autoclean && apt-get update --fix-missing && apt-get upgrade -y && apt-get install curl build-essential -y
#RUN apk update && update-ca-certificates
RUN go clean --modcache
COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GO_GC=off GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ${MAIN_PATH}

## Deploy
FROM scratch AS final
#FROM alpine:3.19.1 AS final
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 80
CMD ["./main"] 