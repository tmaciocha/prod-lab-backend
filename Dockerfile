# Build stage
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /src/app /app/app
EXPOSE 8080
ENTRYPOINT ["/app/app"]