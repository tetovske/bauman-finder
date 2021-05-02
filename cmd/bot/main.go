package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/tetovske/enrollments-parser/config"
	"github.com/tetovske/enrollments-parser/pkg/bot"
	"log"
)

func main() {
	config.Init()

	botApi, err := tgbotapi.NewBotAPI(viper.GetString("tgbot_token")); if err != nil {
		log.Fatal(err)
	}

	bfBot := bot.NewBot(botApi)
	err = bfBot.Start(); if err != nil {
		log.Fatal(err)
	}
}
