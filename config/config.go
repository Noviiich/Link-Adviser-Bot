package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TgBotToken            string
	MongoConnectionString string
}

func MustLoad() Config {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	if tgBotToken == "" {
		log.Fatal("TG_BOT_TOKEN is not specified")
	}
	if mongoConnectionString == "" {
		log.Fatal("MONGO_CONNECTION_STRING is not specified")
	}

	return Config{
		TgBotToken:            tgBotToken,
		MongoConnectionString: mongoConnectionString,
	}
}
