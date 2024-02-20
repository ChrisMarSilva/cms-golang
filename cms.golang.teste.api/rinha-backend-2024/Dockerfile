FROM golang:1.22.0-bullseye as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
ENV GO111MODULE=on
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o app
#RUN go build -o app

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]

#------


# ## Build

# # Stage 1: Build stage
# #FROM golang:1.22 AS builder
# FROM golang:latest AS builder

# # Set the working directory
# WORKDIR /app

# # Set the timezone and install CA certificates
# #RUN apk --no-cache add ca-certificates tzdata
# RUN apk --no-cache add git
# RUN go env -w GOPROXY=direct GOFLAGS="-insecure"
# RUN apt-get update && apt-get install -y ca-certificates openssl
# ARG cert_location=/usr/local/share/ca-certificates
# RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# RUN update-ca-certificates

# # Copy and download dependencies
# RUN go clean --modcache
# COPY go.mod go.sum ./
# #RUN go mod download && go mod verify

# # Copy the source code
# COPY . .

# # Build the Go application
# #RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ${MAIN_PATH}
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ${MAIN_PATH}

# ## Deploy

# # Stage 2: Final stage
# FROM scratch

# RUN apk --no-cache add ca-certificates
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# # Set the working directory
# WORKDIR /app

# # Copy the binary from the build stage
# COPY --from=builder /app/main .

# # Set the entrypoint command
# CMD ["./main"] 

# ----------------------------------

# docker-compose down
# docker-compose up -d --build
# docker run -it rinha-backend-2024-api01:latest
# docker rm -f $(docker ps -a -q)

# ----------------------------------