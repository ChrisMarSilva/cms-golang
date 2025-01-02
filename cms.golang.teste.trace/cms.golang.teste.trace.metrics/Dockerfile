FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /goapp
EXPOSE 5000
CMD ["/goapp"]

# FROM golang:1.23 as builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# # RUN go build -o /goapp
# # RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp ${MAIN_PATH}
# # RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goapp ./main.go
# # RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /goapp ./main.go
# # RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api /*.go

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
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /app/api .
# EXPOSE 8080
# CMD ["./api"]


# FROM golang:1.23.4-bookworm AS build
# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download && go mod verify
# COPY . .
# RUN go generate
# RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GOAMD64=v2 go build -trimpath -tags osusergo,netgo -o server -a -ldflags="-s -w -buildid=" -gcflags="all=-m=0 -l=2 -dwarf=false" -installsuffix cgo

# FROM scratch
# COPY --from=build /app/server /server
# ENTRYPOINT ["/server"]



# FROM golang:1.23.4-bookworm AS build
# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download && go mod verify
# COPY . .
# RUN go generate
# RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux GOAMD64=v2 go build -trimpath -tags osusergo,netgo -o server -a -ldflags="-s -w -buildid=" -gcflags="all=-m=0 -l=2 -dwarf=false" -installsuffix cgo
# FROM scratch
# COPY --from=build /app/server /server
# ENTRYPOINT ["/server"]