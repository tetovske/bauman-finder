package strategies

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type HandleStrategy interface {
	Handle(upd *tgbotapi.Update)
}
