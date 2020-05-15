package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var BotFB *tgbotapi.BotAPI
var err error

func ConnectBotAPI() {
	BotFB, err = tgbotapi.NewBotAPI("917129105:AAGKp2PehTK4b9oEYY3ZR8fY_z2YVqs-qEU")
	if err != nil {
		log.Panic(err)
	}
	BotFB.Debug = true
	log.Printf("Authorized on account %s", BotFB.Self.UserName)

	_, err = BotFB.SetWebhook(tgbotapi.NewWebhook("https://fb-crawler-oaov.herokuapp.com/" + BotFB.Token))
	if err != nil {
		log.Panic(err)
	}

	UpdateMessage()
}
