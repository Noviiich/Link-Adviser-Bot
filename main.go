package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	t := mustToken()
	fmt.Println(t)
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
