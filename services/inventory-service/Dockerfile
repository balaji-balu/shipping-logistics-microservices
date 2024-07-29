FROM golang:1.22 as builder
WORKDIR /app
COPY inventory /app
# Download necessary Go modules (if you have a go.mod file)
COPY inventory/go.mod /app/
RUN go mod download
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o inventory-service 


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/inventory-service .
# Ensure the binary has execution permissions
# RUN chmod +x ./inventory-service

# Debug: List files to verify existence and permissions
#RUN ls -l ./inventory-service
EXPOSE 1323
CMD ["./inventory-service"]