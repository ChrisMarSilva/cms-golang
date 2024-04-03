## Build
#FROM golang:1.22.1 AS builder
FROM golang:1.22.1-alpine3.19 AS builder
WORKDIR /app
#RUN apt-get autoremove && apt-get autoclean && apt-get update --fix-missing && apt-get upgrade -y && apt-get install curl build-essential -y
RUN go clean --modcache
COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GO_GC=off GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

## Deploy
FROM scratch AS final
WORKDIR /app
COPY --from=builder /app/main .
RUN apk --no-cache add sqlite
EXPOSE 8080
CMD ["./main"] 
# 7,66mb

# FROM golang:1.22.1-alpine3.19 AS builder
# WORKDIR /build
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -o app

# FROM alpine:3.19.1 AS final
# RUN adduser --uid 1000 --disabled-password appuser
# USER appuser
# WORKDIR /app
# COPY --from=builder --chown=appuser:appuser /build/app .
# CMD ["./app"] 
# EXPOSE 8080

# 15,05mb
