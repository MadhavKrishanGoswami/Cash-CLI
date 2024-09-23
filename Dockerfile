
# Build Stage
FROM golang:1.23.1-alpine AS builder
WORKDIR /app
COPY . .

# Ensure the Go binary is statically linked
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cash-cli .

# Final Stage: Use scratch as the final image
FROM scratch
WORKDIR /root/
COPY --from=builder /app/cash-cli .

# Set the terminal for color support (if necessary)
ENV TERM=xterm-256color

# Command to run the binary
CMD ["./cash-cli"]
