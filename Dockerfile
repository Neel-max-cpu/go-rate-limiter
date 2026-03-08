# Use official Go image
FROM golang:1.25

# Set working directory inside container
WORKDIR /app

# Copy go mod files first (Copies dependency files first.)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy rest of project
COPY . .

# Build Go binary
RUN go build -o server cmd/server/main.go

# Expose API port
EXPOSE 8080

# Start server
CMD ["./server"]