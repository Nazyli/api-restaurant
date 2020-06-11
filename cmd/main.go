package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nazyli/api-restaurant/delivery/api"
	_categoryMysql "github.com/nazyli/api-restaurant/domain/category/mysql"
	_customerMysql "github.com/nazyli/api-restaurant/domain/customer/mysql"
	_employeeMysql "github.com/nazyli/api-restaurant/domain/employee/mysql"
	_menuMysql "github.com/nazyli/api-restaurant/domain/menu/mysql"
	_positionMysql "github.com/nazyli/api-restaurant/domain/position/mysql"
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

	// Position
	positionMysql := _positionMysql.New(db)
	log.Println("Position mysql is successfully initialized")

	// Category
	categoryMysql := _categoryMysql.New(db)
	log.Println("Category mysql is successfully initialized")

	// Menu
	menuMysql := _menuMysql.New(db)
	log.Println("Menu mysql is successfully initialized")

	// Customer
	customerMysql := _customerMysql.New(db)
	log.Println("Customer mysql is successfully initialized")

	// Employee
	employeeMysql := _employeeMysql.New(db)
	log.Println("Employee mysql is successfully initialized")

	service := service.New(cfg.APP_ID, userMysql, positionMysql, categoryMysql, menuMysql, customerMysql, employeeMysql)
	router := mux.NewRouter()
	api.New(cfg.CDNClaudinary, service).Register(router)
	log.Println("API successfully initialized")

	log.Println("Webserver succesfully started")
	log.Println("Listening to port ", cfg.Webserver)
	log.Fatal(http.ListenAndServe(cfg.Webserver, router))
}
