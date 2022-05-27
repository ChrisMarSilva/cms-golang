
FROM golang:1.18 AS builder
WORKDIR /app
RUN go env -w GOPROXY=direct GOFLAGS="-insecure"
RUN apt-get update && apt-get install -y ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
RUN update-ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main ${MAIN_PATH}

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/.env .
CMD ["./main"] 

# FROM golang:1.18 AS builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . ./
# RUN go build -o /server
#
# FROM gcr.io/distroless/base-debian10:latest   
# WORKDIR /
# COPY --from=builder /server /server
# USER nonroot:nonroot
# ENTRYPOINT ["/server"] 


# FROM golang:1.18.0-alpine3.14 AS builder
# RUN apk add --no-cache --update git
# WORKDIR /app
# COPY . .
# #RUN go get -d -v && go install -v && go build -o /server
# RUN go build -o /server
#
# FROM alpine:3.8
# WORKDIR /
# COPY --from=builder /server /server
# ENTRYPOINT ["/server"] 


# FROM golang:1.18 as builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . ./
# RUN go build -o /server
# ENTRYPOINT ./server
#
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /server /app/server
# ENTRYPOINT ["/server"] 
