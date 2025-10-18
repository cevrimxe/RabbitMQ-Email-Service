# RabbitMQ Email Service

## Proje Açıklaması
RabbitMQ kullanarak email gönderim sistemi. Publisher servis mesajları kuyruğa gönderir, consumer servis kuyruğu dinleyip email gönderir.

## Servisler
1. **Publisher Service**: HTTP API ile mesaj kuyruğa gönderir
2. **Email Consumer Service**: Kuyruğu dinleyip email gönderir
3. **RabbitMQ**: Message broker

## Teknolojiler
- Go 1.23.3
- RabbitMQ
- Docker & Docker Compose
- SMTP (email gönderimi için)

## API Endpoints
### Publisher Service
- `POST /send-email` - Email gönderim isteği

## Mesaj Formatı
```json
{
  "to": "user@example.com",
  "subject": "Konu",
  "body": "Mesaj içeriği"
}
```
