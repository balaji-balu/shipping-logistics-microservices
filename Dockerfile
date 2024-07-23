FROM golang:1.22 as builder
WORKDIR /app
COPY order /app
# Download necessary Go modules (if you have a go.mod file)
COPY order/go.mod /app/
RUN go mod download
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service 


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/order-service .
# Ensure the binary has execution permissions
# RUN chmod +x ./order-service

# Debug: List files to verify existence and permissions
#RUN ls -l ./order-service
EXPOSE 1323
CMD ["./order-service"]