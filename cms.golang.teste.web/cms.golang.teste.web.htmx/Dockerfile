# Build.
FROM golang:1.22.1 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --chown=nonroot --from=build-stage /entrypoint /entrypoint
COPY --chown=nonroot --from=build-stage /app/assets /assets
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]


# FROM golang:1.20.4 AS build
# WORKDIR /go/src/app
# COPY . .
# ENV CGO_ENABLED=0 GOOS=linux GOPROXY=direct
# RUN go build -v -o app .

# FROM scratch
# COPY --from=build /go/src/app/app /go/bin/app
# ENTRYPOINT ["/go/bin/app"]