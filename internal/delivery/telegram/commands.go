package telegram

import (
	"shop-bot/internal/delivery/telegram/buttons"
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	lang := utils.GetLanguage(update, config)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "welcome_message"))

	msg.ReplyMarkup = buttons.MainMenuKeyboard(lang)
	_, err := bot.Send(msg)
	return err
}

func HelpCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	lang := utils.GetLanguage(update, config)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "help_message"))
	_, err := bot.Send(msg)
	return err
}
