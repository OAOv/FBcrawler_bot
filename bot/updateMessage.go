package bot

import (
	"log"
	"net/http"
)

func UpdateMessage() {
	updates := BotFB.ListenForWebhook("/" + BotFB.Token)
	go http.ListenAndServe("0.0.0.0:5000", nil)

	/*
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := BotFB.GetUpdatesChan(u)
		if err != nil {
			log.Panic(err)
		}
	*/

	for update := range updates {
		log.Printf("%+v\n", update)
		/*isCommand := update.Message.IsCommand()

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
		}*/
	}
}
