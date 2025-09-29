# Get Go image
FROM golang:1.25.1 AS build-stage

# Create Working directory
WORKDIR /app

# Copy requirement
COPY go.mod go.sum ./

# Download requirement
RUN go mod download

# Copy application
COPY main.go ./

# Compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o /storage-app

# Get a light image
FROM alpine:3.22 AS run-stage

# Copy compiled app
COPY --from=build-stage /storage-app /storage-app

# Expose port
EXPOSE 8080

# Run the app
CMD ["/storage-app"]