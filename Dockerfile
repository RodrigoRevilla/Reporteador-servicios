# Build the app
FROM golang:alpine AS builder

# Install necessary packages
RUN apk update && apk add --no-cache git

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Copy assets explicitly
COPY assets /build/assets 

# Build the application
RUN go build -o main ./cmd/...

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist
RUN cp /build/main .

# Build a small final image
FROM scratch

# Copy assets and binary
COPY --from=builder /build/assets /assets
COPY --from=builder /dist/main /
COPY --from=builder /build/pkg/api/openapi/api.yml /

# Command to run
ENTRYPOINT ["/main"]
