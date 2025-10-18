# RabbitMQ Email Servisi

[🇹🇷 Türkçe](README.tr.md) | [🇺🇸 English](README.md)

RabbitMQ message broker kullanarak Go ile geliştirilmiş mikroservis tabanlı email sistemi.

## 🏗️ Mimari

- **Publisher Servisi**: Email isteklerini alıp RabbitMQ'ya gönderen HTTP API
- **Consumer Servisi**: Kuyruktan mesajları işleyip SMTP ile email gönderen servis
- **RabbitMQ**: Asenkron iletişim için message broker
- **Docker**: Health check'ler ile konteynerize deployment

## 🚀 Hızlı Başlangıç

1. Projeyi klonla
2. Environment değişkenlerini ayarla:
   ```bash
   cp env.example .env
   # .env dosyasını SMTP bilgilerin ile düzenle
   ```

3. Docker Compose ile servisleri başlat:
   ```bash
   docker-compose up --build
   ```

## 📡 API Kullanımı

Email göndermek için:
```bash
curl -X POST http://localhost:8080/send-email \
  -H "Content-Type: application/json" \
  -d '{
    "to": "user@example.com",
    "subject": "Test Email",
    "body": "Bu bir test email mesajıdır."
  }'
```

## 🔧 Servisler

| Servis | Port | Açıklama |
|--------|------|----------|
| **Publisher** | 8080 | Email istekleri için HTTP API |
| **Consumer** | - | Arka planda email işleyici |
| **RabbitMQ** | 5672 | AMQP protokolü |
| **RabbitMQ Management** | 15672 | Web arayüzü (guest/guest) |

## 🐰 RabbitMQ Yönetimi

Yönetim arayüzüne http://localhost:15672 adresinden erişebilirsiniz
- Kullanıcı adı: `guest`
- Şifre: `guest`

## 🛠️ Environment Değişkenleri

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

## 📋 Özellikler

- ✅ Asenkron email işleme
- ✅ Mikroservis mimarisi
- ✅ Docker konteynerizasyonu
- ✅ Health check'ler ve retry logic
- ✅ RabbitMQ mesaj dayanıklılığı
- ✅ SMTP email teslimatı
- ✅ JSON API arayüzü

## 🔄 Mesaj Akışı

1. **HTTP İsteği** → Publisher Servisi
2. **JSON Mesajı** → RabbitMQ Kuyruğu
3. **Kuyruk İşleme** → Consumer Servisi
4. **SMTP Teslimatı** → Email Gönderildi

## 🏃‍♂️ Geliştirme

Her servisin bağımsız geliştirme için kendi Go modülü var:

```bash
# Publisher servisi
cd publisher && go mod tidy

# Consumer servisi  
cd consumer && go mod tidy
```

## 📊 İzleme

- Logları görüntüle: `docker-compose logs -f consumer`
- RabbitMQ metrikleri: http://localhost:15672
- Kuyruk durumu ve mesaj sayıları yönetim arayüzünde mevcut

## 🤝 Katkıda Bulunma

1. Repository'yi fork et
2. Feature branch oluştur
3. Değişikliklerini yap
4. Pull request gönder
