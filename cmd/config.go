package main

import (
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
}

func loadConfig() (*config, error) {
	var cfg config
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
		return nil, err
	}
	database := conn.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	claudinary := api.CloudinaryConfig{
		AccountName: os.Getenv("CLOUDINARY_NAME"),
		APIKey:      os.Getenv("CLOUDINARY_API_KEY"),
		APISecret:   os.Getenv("CLOUDINARY_API_SECRET"),
	}
	webserver := os.Getenv("WEBSERVER_LISTEN_ADDRESS")
	app_id, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 8)
	if err != nil {
		log.Fatalf("Error Generated APP ID")
		return nil, err
	}
	cfg = config{
		APP_ID:        app_id,
		CDNClaudinary: claudinary,
		Database:      database,
		Webserver:     webserver,
	}
	return &cfg, nil
}