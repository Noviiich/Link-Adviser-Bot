package event_consumer

import (
	"context"
	"log"
	"time"

	"github.com/Noviiich/Link-Adviser-Bot/events"
)

type Consumer struct {
	processor events.Processor
	fetcher   events.Fetcher
	batchSize int
}

func New(f events.Fetcher, p events.Processor, batchSize int) *Consumer {
	return &Consumer{
		fetcher:   f,
		processor: p,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(context.Background(), c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		//организация параллельной обработки, исп sync.WaurtGroup{}
		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

func (c Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(context.Background(), event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
