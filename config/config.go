package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var (
	instance *Configuration
	once     sync.Once
)

// Configuration ...
type Configuration struct {
	TelegramAPITOKEN string
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	ServerPort       int
	ServerHost       string
}

// load ...
func load() *Configuration {
	return &Configuration{
		ServerHost:       cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		TelegramAPITOKEN: cast.ToString(getOrReturnDefault("TELEGRAM_TOKEN", "Not TOKEN")),
		ServerPort:       cast.ToInt(getOrReturnDefault("SERVER_PORT", "8000")),
		PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDB:       cast.ToString(getOrReturnDefault("POSTGRES_DB", "")),
		PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRES_USER", "")),
		PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

//Config ...
func Config() *Configuration {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	once.Do(func() {
		instance = load()
	})

	return instance
}
