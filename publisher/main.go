package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewPublisher(rabbitmqURL string) (*Publisher, error) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// Kuyruk oluştur
	_, err = ch.QueueDeclare(
		"email_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		conn:    conn,
		channel: ch,
	}, nil
}

func (p *Publisher) PublishEmail(email EmailRequest) error {
	body, err := json.Marshal(email)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		"",            // exchange
		"email_queue", // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}

func (p *Publisher) Close() {
	p.channel.Close()
	p.conn.Close()
}

func main() {
	// RabbitMQ bağlantısını bekle
	var publisher *Publisher
	var err error

	for i := 0; i < 30; i++ {
		publisher, err = NewPublisher("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
		log.Printf("RabbitMQ bağlantısı bekleniyor... (%d/30)", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("RabbitMQ bağlantısı başarısız:", err)
	}
	defer publisher.Close()

	r := mux.NewRouter()

	r.HandleFunc("/send-email", func(w http.ResponseWriter, r *http.Request) {
		var emailReq EmailRequest
		if err := json.NewDecoder(r.Body).Decode(&emailReq); err != nil {
			http.Error(w, "Geçersiz JSON", http.StatusBadRequest)
			return
		}

		if err := publisher.PublishEmail(emailReq); err != nil {
			http.Error(w, "Email kuyruğa gönderilemedi", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email kuyruğa gönderildi"})
	}).Methods("POST")

	log.Println("Publisher service 8080 portunda başlatılıyor...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
