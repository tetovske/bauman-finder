package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tetovske/enrollments-parser/pkg/bot/strategies"
	"log"
)

type Bot struct {
	botApi *tgbotapi.BotAPI
}

func NewBot(botApi *tgbotapi.BotAPI) *Bot {
	return &Bot{botApi: botApi}
}

func (bot *Bot) Start() error {
	updates, err := bot.initializeUpdates()
	if err != nil {
		return err
	}

	bot.handleUpdates(updates)

	return err
}

func (bot *Bot) initializeUpdates() (tgbotapi.UpdatesChannel, error) {
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60

	updates, err := bot.botApi.GetUpdatesChan(upd)

	return updates, err
}

func (bot *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] Message: %s", update.Message.From.UserName, update.Message.Text)

		err := bot.handleUpdate(update)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (bot *Bot) handleUpdate(upd tgbotapi.Update) error {
	resp, err := strategies.StrategyFactory(&upd).Handle(&upd)
	if err != nil {
		return err
	}

	err = resp.SendMessage(bot.botApi)

	return err
}
