## Build
FROM golang:1.24.5-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify
COPY . . 
#RUN GO111MODULE=on CGO_ENABLED=0 GO_GC=off GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ${MAIN_PATH}
RUN CGO_ENABLED=0 GO_GC=off GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o app ./cmd/api/main.go

## Deploy
FROM scratch AS final
#RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /src/app .
#RUN adduser -D -s /bin/sh appuser
#USER appuser
EXPOSE 8080
CMD ["./app"]