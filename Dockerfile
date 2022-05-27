# Start from golang base image
FROM golang:1.18-buster as builder

WORKDIR /app

COPY go.* ./

RUN go mod download 

COPY . ./

# Build the Go app
RUN go build -o server .

FROM debian:buster-slim

WORKDIR /app

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/.env /app/
COPY --from=builder /app/log/. /app/log/

ENV ENVIRONMENT docker

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/app/server" ]