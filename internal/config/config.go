package config

import "os"

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func Load() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "root:YOUR_PASSWORD@tcp(127.0.0.1:3306)/workdirk?parseTime=true"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	return &Config{
		DatabaseURL: dbURL,
		ServerPort:  port,
	}
}
