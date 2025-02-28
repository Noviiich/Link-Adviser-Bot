package main

import (
	"flag"
	"log"

	tgClient "github.com/Noviiich/Link-Adviser-Bot/clients/telegram"
	event_consumer "github.com/Noviiich/Link-Adviser-Bot/consumer/event-consumer"
	"github.com/Noviiich/Link-Adviser-Bot/events/telegram"
	"github.com/Noviiich/Link-Adviser-Bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eProcessor, eProcessor, batchSize)
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
