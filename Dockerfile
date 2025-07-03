FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN GOOS=linux go build -o /main .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /main .
COPY .config ./.config/
COPY images ./images/

ENV APP_ENV=prod
	
CMD ["./main"]
