package settings_handler

import (
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleChangeLanguage(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	lang := utils.GetLanguage(update, config)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "select_language"))
	msg.ReplyMarkup = utils.LanguageKeyboard(config.AllowedLanguages)

	_, err := bot.Send(msg)
	return err
}

func SaveLanguagePreference(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	newLang := update.CallbackQuery.Data // Новый язык из callback data
	chatID := update.CallbackQuery.Message.Chat.ID

	if !utils.LanguageIsAllowed(newLang, config.AllowedLanguages) {
		msg := tgbotapi.NewMessage(chatID, "Этот язык недоступен.")
		_, err := bot.Send(msg)
		return err
	}

	// Сохранение языка в базу данных или кэш
	err := utils.SaveUserLanguage(chatID, newLang)
	if err != nil {
		return err
	}

	// Отправляем подтверждение
	msg := tgbotapi.NewMessage(chatID, utils.Translate(newLang, "language_changed"))
	_, err = bot.Send(msg)
	return err
}
