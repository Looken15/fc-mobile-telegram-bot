FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

FROM alpine:latest
WORKDIR /app

# Копируем бинарник
COPY --from=builder /main .

# Копируем папку .config в рабочую директорию /app
COPY .config ./.config/

# Проверяем что файлы скопировались (для отладки)
RUN ls -la ./.config && \
    [ -f ./.config/dev.json ] && echo "dev.json exists" || echo "dev.json missing"
	
CMD ["./main"]
