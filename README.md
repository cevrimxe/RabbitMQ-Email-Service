# RabbitMQ Email Service

[🇹🇷 Türkçe](README.tr.md) | [🇺🇸 English](README.md)

A microservice-based email system using RabbitMQ message broker with Go services.

## 🏗️ Architecture

- **Publisher Service**: HTTP API that receives email requests and publishes to RabbitMQ
- **Consumer Service**: Processes messages from queue and sends emails via SMTP
- **RabbitMQ**: Message broker for asynchronous communication
- **Docker**: Containerized deployment with health checks

## 🚀 Quick Start

1. Clone the repository
2. Set up environment variables:
   ```bash
   cp env.example .env
   # Edit .env file with your SMTP credentials
   ```

3. Start services with Docker Compose:
   ```bash
   docker-compose up --build
   ```

## 📡 API Usage

Send an email:
```bash
curl -X POST http://localhost:8080/send-email \
  -H "Content-Type: application/json" \
  -d '{
    "to": "user@example.com",
    "subject": "Test Email",
    "body": "This is a test email message."
  }'
```

## 🔧 Services

| Service | Port | Description |
|---------|------|-------------|
| **Publisher** | 8080 | HTTP API for email requests |
| **Consumer** | - | Background email processor |
| **RabbitMQ** | 5672 | AMQP protocol |
| **RabbitMQ Management** | 15672 | Web UI (guest/guest) |

## 🐰 RabbitMQ Management

Access the management interface at http://localhost:15672
- Username: `guest`
- Password: `guest`

## 🛠️ Environment Variables

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

## 📋 Features

- ✅ Asynchronous email processing
- ✅ Microservice architecture
- ✅ Docker containerization
- ✅ Health checks and retry logic
- ✅ RabbitMQ message durability
- ✅ SMTP email delivery
- ✅ JSON API interface

## 🔄 Message Flow

1. **HTTP Request** → Publisher Service
2. **JSON Message** → RabbitMQ Queue
3. **Queue Processing** → Consumer Service
4. **SMTP Delivery** → Email Sent

## 🏃‍♂️ Development

Each service has its own Go module for independent development:

```bash
# Publisher service
cd publisher && go mod tidy

# Consumer service  
cd consumer && go mod tidy
```

## 📊 Monitoring

- View logs: `docker-compose logs -f consumer`
- RabbitMQ metrics: http://localhost:15672
- Queue status and message counts available in management UI

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request