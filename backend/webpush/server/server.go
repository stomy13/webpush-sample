package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MasatoTokuse/webpush/webpush/dbaccess"
	myjson "github.com/MasatoTokuse/webpush/webpush/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type ConnectArgs struct {
	Address  string
	Port     string
	DBName   string
	User     string
	Password string
}

type Serve interface {
	RunServer(port string, conargs *ConnectArgs) error
}
type server struct{}

var Server *server

func NewServer() *server {
	return Server
}

func (*server) RunServer(port string, conargs *ConnectArgs) error {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// response public key
	r.Get("/pubkey", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("BO-UY2C7nObUfD6MYDfw5ecSpIuf8REJsu9gISnsCCtdvC6u-FpHkC_HNjjZmjvnn1HzOiGaLJy-tzPfY6M_6ns"))
	})

	// insert subscription
	r.Post("/subscription", func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))

		var jep myjson.Endpoint
		json.Unmarshal(body, &jep)
		// log.Println("-----------unmarshal--------------")
		// log.Println(jep.Endpoint)
		// log.Println(jep.Keys.Key)
		// log.Println(jep.Keys.Token)

		var ep dbaccess.Endpoint
		ep.Endpoint = jep.Endpoint
		ep.Key = jep.Keys.Key
		ep.Token = jep.Keys.Token
		ep.UserID = 50

		// TODO:DBに保存
		db := dbaccess.ConnectGorm()
		defer db.Close()
		db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&dbaccess.Endpoint{})
		db.NewRecord(ep)
		db.Create(&ep)

		w.Write([]byte("ok"))
	})

	return http.ListenAndServe(port, r)
}
