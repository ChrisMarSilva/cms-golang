## Build
FROM golang:1.21-alpine3.18 as builder
#FROM golang:1.22.0 AS builder
RUN apk update 
WORKDIR /src
#WORKDIR /app
#RUN go clean --modcache
#RUN apt-get autoremove && apt-get autoclean && apt-get update --fix-missing && apt-get upgrade -y && apt-get install curl build-essential -y
COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify
COPY . . 
RUN go build -o rinha ./cmd/
#RUN GO111MODULE=on CGO_ENABLED=0 GO_GC=off GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ${MAIN_PATH}

## Deploy
FROM scratch AS final
#FROM alpine:3.18 as final
#WORKDIR /app
COPY --from=builder /src/rinha .
#COPY --from=builder /app/main .
RUN mkdir /pprof
EXPOSE 3000
CMD ["./rinha"]
#CMD [ "app" ]
#CMD ["./main"]