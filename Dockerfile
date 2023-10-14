# Stage 1: Go build environment
FROM golang:1.16-alpine AS build-env

# Set the working directory to the app directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the Go functions
COPY . .
RUN go build -o /app/ExecuteFuzzySearch functions/backend/execute_fuzzy_search.go
RUN go build -o /app/ManageOutboundRouting functions/backend/manage_outbound_routing.go

# Stage 2: Barebones runtime environment
FROM alpine:latest

# Copy the binary files from the first stage
COPY --from=build-env /app/ExecuteFuzzySearch /app/ManageOutboundRouting /app/

# Set the working directory to the app directory
WORKDIR /app

# Expose port 8080 for the OpenFaaS Gateway
EXPOSE 8080

# Start the OpenFaaS function using the appropriate binary file
CMD ["./ManageOutboundRouting"]