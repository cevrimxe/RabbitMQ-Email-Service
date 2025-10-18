package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
)

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailService struct {
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
}

func NewEmailService() *EmailService {
	port, _ := strconv.Atoi(getEnv("SMTP_PORT", "587"))
	return &EmailService{
		smtpHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		smtpPort:     port,
		smtpUsername: getEnv("SMTP_USERNAME", ""),
		smtpPassword: getEnv("SMTP_PASSWORD", ""),
	}
}

func (e *EmailService) SendEmail(msg EmailMessage) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.smtpUsername)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody("text/html", msg.Body)

	d := gomail.NewDialer(e.smtpHost, e.smtpPort, e.smtpUsername, e.smtpPassword)

	return d.DialAndSend(m)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	// RabbitMQ bağlantısını bekle
	var conn *amqp.Connection
	var err error

	for i := 0; i < 30; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
		log.Printf("RabbitMQ bağlantısı bekleniyor... (%d/30)", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("RabbitMQ bağlantısı başarısız:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Kanal açılamadı:", err)
	}
	defer ch.Close()

	// Kuyruk oluştur
	q, err := ch.QueueDeclare(
		"email_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatal("Kuyruk oluşturulamadı:", err)
	}

	// Email servisi
	emailService := NewEmailService()

	// Mesajları tüket
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("Mesaj tüketimi başarısız:", err)
	}

	log.Println("Email consumer servisi başlatıldı. Mesajlar bekleniyor...")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var emailMsg EmailMessage
			if err := json.Unmarshal(d.Body, &emailMsg); err != nil {
				log.Printf("JSON parse hatası: %v", err)
				continue
			}

			log.Printf("Email gönderiliyor: %s -> %s", emailMsg.Subject, emailMsg.To)

			if err := emailService.SendEmail(emailMsg); err != nil {
				log.Printf("Email gönderim hatası: %v", err)
			} else {
				log.Printf("Email başarıyla gönderildi: %s", emailMsg.To)
			}
		}
	}()

	<-forever
}
