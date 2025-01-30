package telegram

import (
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleContacts(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	lang := utils.GetLanguage(update, config)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "contacts_info"))
	_, err := bot.Send(msg)
	return err
}
