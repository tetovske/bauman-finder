package strategies

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tetovske/enrollments-parser/pkg/bot/response"
)

type HandleStrategy interface {
	Handle(upd *tgbotapi.Update) (response.BotResponse, error)
}
