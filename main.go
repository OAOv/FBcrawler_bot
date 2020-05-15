package main

import (
	"FBcrawler/bot"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	bot.ConnectBotAPI()
	router.POST("/"+bot.BotFB.Token, bot.UpdateMessage)

	err := router.Run()
	if err != nil {
		log.Println(err)
	}
}
