package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/Noviiich/Link-Adviser-Bot/clients/telegram"
	event_consumer "github.com/Noviiich/Link-Adviser-Bot/consumer/event-consumer"
	"github.com/Noviiich/Link-Adviser-Bot/events/telegram"
	"github.com/Noviiich/Link-Adviser-Bot/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "files_storage"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
