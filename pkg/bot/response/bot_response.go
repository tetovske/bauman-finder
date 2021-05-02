package response

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type TelegramAPISender interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type BotResponse interface {
	SendMessage(bot TelegramAPISender) error
	GetMessage() *tgbotapi.MessageConfig
}
