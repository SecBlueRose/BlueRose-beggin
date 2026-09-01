# =========================================
# ЭТАП I: Builder (Сборка бинарника)
# =========================================
FROM golang:1.22-alpine AS builder

# Настраиваем рабочую директорию
WORKDIR /app

# Инициализируем Go Module (если еще не создан)
RUN go mod init core-api || true

# Копируем исходный код
COPY main.go .

# Компилируем статический бинарник без CGO и отладочных символов
RUN CGO_enabled=0 GOOS=linux go build -ldflags="-s -w" -o core-api main.go

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