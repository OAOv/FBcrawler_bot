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

	BotFB.RemoveWebhook()

	_, err = BotFB.SetWebhook(tgbotapi.NewWebhookWithCert("https://fb-crawler-oaov.herokuapp.com/"+BotFB.Token, nil))
	if err != nil {
		log.Panic(err)
	}

	info, err := BotFB.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}

	UpdateMessage()
}
