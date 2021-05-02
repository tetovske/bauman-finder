package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/tetovske/enrollments-parser/pkg/bot/response"
	"github.com/tetovske/enrollments-parser/pkg/bot/strategies"
)

type CommandStrategy interface {
	strategies.HandleStrategy
	getCommandName() string
}

type StartStrategy struct {}

func NewStartStrategy() *StartStrategy {
	return &StartStrategy{}
}

func (s *StartStrategy) Handle(update *tgbotapi.Update) (response.BotResponse, error) {
	return response.NewInstantResponse(
		update.Message.Chat.ID,
		viper.GetString("commands.start"),
		response.TextMessage), nil
}


