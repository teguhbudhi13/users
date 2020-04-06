package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ServerConfig for server configuration
type ServerConfig struct {
	Port           string
	AllowedOrigins string
}

// Config struct
type Config struct {
	DB     *DBConfig
	Server *ServerConfig
}

// DBConfig struct define config
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig configur url
func GetConfig() *Config {
	// load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error port number")
	}

	return &Config{
		DB: &DBConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_DATABASE"),
			Charset:  "utf8",
		},
		Server: &ServerConfig{
			Port:           os.Getenv("SERVER_PORT"),
			AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),
		},
	}
}
