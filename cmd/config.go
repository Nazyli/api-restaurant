package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nazyli/api-restaurant/delivery/api"
	conn "github.com/nazyli/api-restaurant/util/database/mysql"
)

type config struct {
	APP_ID        int64
	CDNClaudinary api.CloudinaryConfig
	Database      conn.DBConfig
	Webserver     string
	Tax           float64
}

func loadConfig() (*config, error) {
	var cfg config
	var err error
	_ = godotenv.Load()
	if os.Getenv("APP_NAME") == "" {
		log.Fatalf("Get App Name not success")
	}
	log.Println("Starting ", os.Getenv("APP_NAME"))
	// if err != nil && strings.Contains(err.Error(), "directory")  {
	// 	log.Fatalf("Error getting env, %v", err)
	// 	return nil, err
	// }
	database := conn.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dialect:  os.Getenv("DB_DIALECT"),
	}
	claudinary := api.CloudinaryConfig{
		AccountName: os.Getenv("CLOUDINARY_NAME"),
		APIKey:      os.Getenv("CLOUDINARY_API_KEY"),
		APISecret:   os.Getenv("CLOUDINARY_API_SECRET"),
	}

	webserver := os.Getenv("WEBSERVER_LISTEN_ADDRESS")
	if os.Getenv("PORT") != "" {
		webserver = fmt.Sprintf(":%v", os.Getenv("PORT"))
	}
	app_id, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 8)
	if err != nil {
		log.Fatalf("Error Generated APP ID")
		return nil, err
	}
	tax, err := strconv.ParseFloat(os.Getenv("TAX"), 10)
	cfg = config{
		APP_ID:        app_id,
		CDNClaudinary: claudinary,
		Database:      database,
		Webserver:     webserver,
		Tax:           tax,
	}
	return &cfg, nil
}
