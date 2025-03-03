package main

import (
	"log"
	"time"

	tgClient "github.com/Noviiich/Link-Adviser-Bot/clients/telegram"
	"github.com/Noviiich/Link-Adviser-Bot/config"
	event_consumer "github.com/Noviiich/Link-Adviser-Bot/consumer/event-consumer"
	"github.com/Noviiich/Link-Adviser-Bot/events/telegram"
	"github.com/Noviiich/Link-Adviser-Bot/storage/mongo"
	_ "github.com/mattn/go-sqlite3"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "files_storage"
	// sqliteStoragePath = "data/sqlite/storage.db"
	batchSize = 100
)

func main() {
	cfg := config.MustLoad()
	//storage := files.New(storagePath)

	storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
