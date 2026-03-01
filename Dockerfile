FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o student-api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/student-api .
EXPOSE 8080
CMD ["./student-api"]
