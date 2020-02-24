package message

import (
	"encoding/json"
	"log"

	"github.com/MasatoTokuse/webpush/webpush/dbaccess"
	"github.com/MasatoTokuse/webpush/webpush/setting"
	webpush "github.com/SherClockHolmes/webpush-go"
)

type message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (msg *message) Push(conargs *dbaccess.ConnectArgs) error {
	// server keypair
	keypair, err := setting.GetKeypair()
	if err != nil {
		return err
	}

	// Select subscription
	db := dbaccess.ConnectGorm(conargs)
	defer db.Close()

	var s dbaccess.Subscription
	db.Where("user_id = ?", 50).Last(&s)

	// Decode subscription
	sub := &webpush.Subscription{Endpoint: s.Endpoint, Keys: webpush.Keys{P256dh: s.P256dh, Auth: s.Auth}}

	messageJSON, _ := json.MarshalIndent(msg, "", "  ")

	// Send Notification
	resp, err := webpush.SendNotification(messageJSON, sub, &webpush.Options{
		Subscriber:      "example@example.com",
		VAPIDPublicKey:  keypair.PublicKey,
		VAPIDPrivateKey: keypair.PrivateKey,
		TTL:             30,
	})
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return err
}

func NewMessage(title, body string) *message {
	return &message{
		Title: title,
		Body:  body,
	}
}
