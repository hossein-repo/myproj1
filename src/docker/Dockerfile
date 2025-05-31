FROM golang:latest

# ساختار کاری
WORKDIR /app

# کپی کل src به مسیر مشخص
COPY src /app/src

# تنظیم محیط کاری واقعی
WORKDIR /app/src

# دانلود ماژول‌ها
RUN go mod download

# کامپایل به مسیر ثابت
RUN go build -o /app/main ./cmd/main.go

# اجرای فایل باینری
CMD ["/app/main"]
