package bot

import (
	"FBcrawler/api"
	"FBcrawler/task"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func UpdateMessage() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := BotFB.ListenForWebhook("/hook")
	go http.ListenAndServe("https://fb-crawler-oaov.herokuapp.com/hook", nil)
	/*updates, err := BotFB.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}*/

	for update := range updates {
		isCommand := update.Message.IsCommand()

		if isCommand {
			userInput := update.Message.Command()
			Target := task.GetTarget(userInput)

			if Target == "NotFound" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter correct command.")
				if _, err := BotFB.Send(msg); err != nil {
					log.Panic(err)
				}
				return
			}

			datas := api.Crawler(Target)

			for _, data := range *datas {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, data.Title+"\n\n"+data.URL)

				if _, err := BotFB.Send(msg); err != nil {
					log.Panic(err)
				}
			}
		} else {
			if update.Message.Text != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

				if _, err := BotFB.Send(msg); err != nil {
					log.Panic(err)
				}
			} else {
				smsg := tgbotapi.NewStickerShare(update.Message.Chat.ID, update.Message.Sticker.FileID)

				if _, err := BotFB.Send(smsg); err != nil {
					log.Panic(err)
				}
			}
		}
	}
}
