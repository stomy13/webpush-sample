package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MasatoTokuse/webpush/webpush/dbaccess"
	"github.com/MasatoTokuse/webpush/webpush/setting"
	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Serve interface {
	RunServer(port string, conargs *dbaccess.ConnectArgs) error
}
type server struct{}

var Server *server

func NewServer() *server {
	return Server
}

func (*server) RunServer(port string, conargs *dbaccess.ConnectArgs) error {

	// server keypair
	keypair, err := setting.GetKeypair()
	if err != nil {
		return err
	}

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// response public key
	r.Get("/pubkey", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(keypair.PublicKey))
	})

	// insert subscription
	r.Post("/subscription", func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))

		var js webpush.Subscription
		json.Unmarshal(body, &js)

		var sub dbaccess.Subscription
		sub.Endpoint = js.Endpoint
		sub.P256dh = js.Keys.P256dh
		sub.Auth = js.Keys.Auth
		sub.UserID = 50

		db := dbaccess.ConnectGorm(conargs)
		defer db.Close()
		db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&dbaccess.Subscription{})
		db.NewRecord(sub)
		db.Create(&sub)

		w.Write([]byte("ok"))
	})

	return http.ListenAndServe(port, r)
}
