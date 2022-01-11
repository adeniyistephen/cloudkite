# Build the Go Binary.
FROM golang:1.17 as build_cloudkite-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /cloudkite

# Build the service binary.
WORKDIR /cloudkite/app
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.15
COPY --from=build_cloudkite-api /cloudkite/app/app /cloudkite/app
WORKDIR /cloudkite
CMD ["./app"]