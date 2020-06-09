package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nazyli/api-restaurant/delivery/api"
	_userMysql "github.com/nazyli/api-restaurant/domain/user/mysql"
	"github.com/nazyli/api-restaurant/service"
	conn "github.com/nazyli/api-restaurant/util/database/mysql"
)

func init() {
	log.SetPrefix("[API-RESATAURANT] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {
	var (
		cfg, err = loadConfig()
	)
	if err != nil {
		log.Println(err)
	}
	db, err := conn.Init(cfg.Database)
	if err != nil {
		panic(err)
	}
	log.Println("Database successfully initialized")

	// User
	userMysql := _userMysql.New(db)
	log.Println("User mysql is successfully initialized")

	service := service.New(userMysql)
	router := mux.NewRouter()
	api.New(cfg.APP_ID, cfg.CDNClaudinary, service).Register(router)
	log.Println("API successfully initialized")

	log.Println("Webserver succesfully started")
	log.Println("Listening to port ", cfg.Webserver)
	log.Fatal(http.ListenAndServe(cfg.Webserver, router))
}
