package settings_handler

import (
	"shop-bot/internal/delivery/telegram/buttons"
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleSettings(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	lang := utils.GetLanguage(update, config)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "settings_menu"))
	msg.ReplyMarkup = buttons.SettingsMenuKeyboard(lang)

	_, err := bot.Send(msg)
	return err
}
