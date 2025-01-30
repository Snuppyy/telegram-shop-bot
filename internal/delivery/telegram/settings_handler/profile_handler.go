package settings_handler

import (
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleProfile(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	lang := utils.GetLanguage(update, config)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "profile_settings"))
	msg.ReplyMarkup = utils.ProfileMenuKeyboard(lang)

	_, err := bot.Send(msg)
	return err
}

func UpdateProfileField(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig, field string) error {
	chatID := update.Message.Chat.ID
	newValue := update.Message.Text

	// Сохраняем пользовательское поле (например, имя) в базу данных
	err := utils.UpdateUserProfile(chatID, field, newValue)
	if err != nil {
		return err
	}

	// Отправляем подтверждение обновления
	lang := utils.GetLanguage(update, config)
	confirmation := utils.Translate(lang, "profile_updated")
	msg := tgbotapi.NewMessage(chatID, confirmation)
	_, err = bot.Send(msg)
	return err
}
