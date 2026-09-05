# =========================================
# ЭТАП I: Builder (Сборка бинарника)
# =========================================
FROM golang:1.22-alpine AS builder

# Настраиваем рабочую директорию
WORKDIR /app
RUN go mod init core-api || true
COPY main.go .

# Add this line to fetch the required dependency
RUN go get github.com/lib/pq
# RUN go mod tidy (Optional: helps clean up dependencies)

# Your existing build command
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o core-api main.go

# =============================================
# ЭТАП II: Runner(Минимальное и безопасное окружение)
# =============================================
FROM alpine:latest AS runner

WORKDIR /app

# Копируем из builder-контейнера только готовый бинарник
COPY --from=builder /app/core-api .

# Указываем порт, который слушает приложение
EXPOSE 8080

# Команда для запуска
CMD ["./core-api"]