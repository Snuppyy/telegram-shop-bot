package telegram

import (
	"errors"
	"log"
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	CommandHandlers map[string]func(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error
	ButtonHandlers  map[string]func(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error
}

func NewRouter() *Router {
	return &Router{
		CommandHandlers: make(map[string]func(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error),
		ButtonHandlers:  make(map[string]func(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error),
	}
}

func (r *Router) RegisterCommand(command string, handler func(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error) {
	r.CommandHandlers[command] = handler
}

func (r *Router) RegisterButtonCallback(callbackData string, handler func(callback *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error) {
	r.ButtonHandlers[callbackData] = handler
}

func (r *Router) Handle(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	if update.Message != nil && update.Message.IsCommand() {
		handler, ok := r.CommandHandlers[update.Message.Command()]
		if ok {
			return handler(update, bot, config)
		}
		log.Printf("Unknown command: %s", update.Message.Command())
		return errors.New("unknown command")
	}

	if update.CallbackQuery != nil {
		handler, ok := r.ButtonHandlers[update.CallbackQuery.Data]
		if ok {
			return handler(update.CallbackQuery, bot, config)
		}
		log.Printf("Unknown button callback: %s", update.CallbackQuery.Data)
		return errors.New("unknown button callback")
	}

	return nil
}
