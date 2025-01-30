package telegram

import (
	"log"
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	Bot    *tgbotapi.BotAPI
	Router *Router
	Config *utils.TelegramConfig
}

func NewBotHandler(config *utils.TelegramConfig, router *Router) (*BotHandler, error) {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, err
	}

	return &BotHandler{
		Bot:    bot,
		Router: router,
		Config: config,
	}, nil
}

func (h *BotHandler) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := h.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		err := h.Router.Handle(update, h.Bot, h.Config)
		if err != nil {
			log.Printf("Error handling update: %v", err)
		}
	}

	return nil
}
