package response

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type MessageType int

const (
	TextMessage MessageType = iota
	AnimatedMessage
)

type InstantResponse struct {
	chatID    int64
	msgType   MessageType
	msgConfig tgbotapi.MessageConfig
}

func (r *InstantResponse) MsgConfig() tgbotapi.MessageConfig {return r.msgConfig}

func (r *InstantResponse) ChatID() int64 {return r.chatID}

func (r *InstantResponse) MsgType() MessageType {return r.msgType}

func NewInstantResponse(chatID int64, msg string, msgType MessageType) *InstantResponse {
	msgConfig := tgbotapi.NewMessage(chatID, msg)

	return &InstantResponse{
		chatID: 	chatID,
		msgType:	msgType,
		msgConfig:	msgConfig,
	}
}

func (r *InstantResponse) GetMessage() *tgbotapi.MessageConfig {
	return &r.msgConfig
}

func (r *InstantResponse) SendMessage(bot TelegramAPISender) error {
	_, err := bot.Send(r.msgConfig); if err != nil {
		log.Fatal(err)
	}

	return err
}
