# RabbitMQ Email Service

RabbitMQ kullanarak email gönderim sistemi.

## Kurulum

1. Projeyi klonla
2. Environment değişkenlerini ayarla:
   ```bash
   cp env.example .env
   # .env dosyasını düzenle
   ```

3. Docker Compose ile başlat:
   ```bash
   docker-compose up --build
   ```

## Kullanım

Email göndermek için:
```bash
curl -X POST http://localhost:8080/send-email \
  -H "Content-Type: application/json" \
  -d '{
    "to": "user@example.com",
    "subject": "Test Email",
    "body": "Bu bir test emailidir."
  }'
```

## Servisler

- **Publisher**: Port 8080 - HTTP API
- **Consumer**: Email gönderim servisi
- **RabbitMQ**: Port 5672 (AMQP), Port 15672 (Management UI)

## RabbitMQ Management

http://localhost:15672 adresinden RabbitMQ yönetim paneline erişebilirsiniz.
- Kullanıcı: guest
- Şifre: guest
