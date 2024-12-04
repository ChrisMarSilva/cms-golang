FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /goapp
EXPOSE 8080
CMD ["/goapp"]

# FROM golang:1.23 as builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o /goapp
# # RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp ${MAIN_PATH}
# # RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp ./main.go
# # RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /goapp ./main.go

# # FROM builder AS tester
# # RUN go test -v ./...

# FROM scratch AS tester
# WORKDIR /app
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# # COPY --from=builder /app/goapp .
# # COPY --from=builder /app /app/goapp
# COPY --from=builder /goapp /goapp
# EXPOSE 8080
# CMD ["/goapp"]