package strategies

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/tetovske/enrollments-parser/pkg/bot/response"
)

type CommandStrategy interface {
	HandleStrategy
	GetCommandName() string
}

type StartStrategy struct {
	CommandStrategy
	commandName string
}

func NewStartStrategy() *StartStrategy {
	return &StartStrategy{}
}

func (s *StartStrategy) GetCommandName() string {
	return s.commandName
}

func (s *StartStrategy) Handle(upd *tgbotapi.Update) (response.BotResponse, error) {
	return response.NewInstantResponse(
		upd.Message.Chat.ID,
		viper.GetString("commands.start"),
		response.TextMessage), nil
}
