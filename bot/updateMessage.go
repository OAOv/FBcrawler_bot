package bot

import (
	"FBcrawler/api"
	"FBcrawler/task"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func UpdateMessage(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		isCommand := update.Message.IsCommand()

		if isCommand {
			userInput := update.Message.Command()
			Target := task.GetTarget(userInput)

			if Target == "NotFound" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter correct command.")
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
				return
			}

			datas := api.Crawler(Target)

			for _, data := range *datas {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, data.Title+"\n\n"+data.URL)

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}
		} else {
			if update.Message.Text != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			} else {
				smsg := tgbotapi.NewStickerShare(update.Message.Chat.ID, update.Message.Sticker.FileID)

				if _, err := bot.Send(smsg); err != nil {
					log.Panic(err)
				}
			}
		}
	}
}
