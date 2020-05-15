package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ConnectBotAPI() {
	bot, err := tgbotapi.NewBotAPI("917129105:AAGKp2PehTK4b9oEYY3ZR8fY_z2YVqs-qEU")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	UpdateMessage(bot)
}
