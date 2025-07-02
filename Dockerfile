FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /main .
COPY .config ./.config/

ENV APP_ENV=prod
ENV TOKEN=${TOKEN}
	
CMD ["./main"]
