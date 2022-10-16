package main

import (
	"context"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/onmono/clean-architecture/internal/config/tinkoff"
	"github.com/onmono/clean-architecture/internal/domain/entity/tinkoff/sandbox"
	tinkoff "github.com/onmono/clean-architecture/pkg/client/tinkoff/sandbox"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()
	cfg := config.NewConfig()
	sandboxClient := tinkoff.NewClient(cfg)

	balance := sandbox.Balance{
		Ticker: "USD", Balance: 5000000,
	}

	sandboxClient.Dial(ctx, balance)
	stream(cfg)
}

func stream(cfg config.TinkoffInvestConfig) {
	logger := log.New(os.Stdout, "[invest-openapi-go-sdk]", log.LstdFlags)

	client, err := sdk.NewStreamingClient(logger, cfg.Token())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	go func() {
		err = client.RunReadLoop(func(event interface{}) error {
			logger.Printf("Got event %+v", event)
			return nil
		})
		if err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Подписка на получение событий по инструменту BBG000N9MNX3 (TSLA)")
	err = client.SubscribeInstrumentInfo("BBG000N9MNX3", requestID())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Подписка на получение свечей по инструменту BBG000N9MNX3 (TSLA)")
	err = client.SubscribeCandle("BBG000N9MNX3", sdk.CandleInterval5Min, requestID())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Подписка на получения стакана по инструменту BBG000N9MNX3 (TSLA)")
	err = client.SubscribeOrderbook("BBG000N9MNX3", 10, requestID())
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(20 * time.Second)

	log.Println("Отписка от получения событий по инструменту BBG000N9MNX3 (TSLA)")
	err = client.UnsubscribeInstrumentInfo("BBG000N9MNX3", requestID())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Отписка от получения свечей по инструменту BBG000N9MNX3 (TSLA)")
	err = client.UnsubscribeCandle("BBG000N9MNX3", sdk.CandleInterval5Min, requestID())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Отписка от получения стакана по инструменту BBG000N9MNX3 (TSLA)")
	err = client.UnsubscribeOrderbook("BBG000N9MNX3", 10, requestID())
	if err != nil {
		log.Fatalln(err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Генерируем уникальный ID для запроса
func requestID() string {
	b := make([]rune, 12)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
