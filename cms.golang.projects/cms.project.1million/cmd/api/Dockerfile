## Build
FROM golang:1.24.6-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify
COPY . . 
RUN CGO_ENABLED=0 GO_GC=off GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o app ./cmd/api/main.go

## Deploy
FROM scratch AS final
WORKDIR /app
COPY --from=builder /src/app .
EXPOSE 8080
CMD ["./app"]