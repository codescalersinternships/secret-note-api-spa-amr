# Build Stage
FROM golang:1.22.4-alpine AS BuildStage

# Install gcc for CGO
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/main.go

# Deploy Stage
FROM alpine:latest AS main
WORKDIR /
COPY .env .env
COPY --from=BuildStage /app/main /main
EXPOSE 8000

ENTRYPOINT [ "/main" ]
