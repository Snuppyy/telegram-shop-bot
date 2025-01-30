package settings_handler

import (
	"shop-bot/internal/delivery/telegram/buttons"
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleAddress(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	lang := utils.GetLanguage(update, config)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "address_menu"))
	msg.ReplyMarkup = buttons.AddressMenuKeyboard(lang)

	_, err := bot.Send(msg)
	return err
}

func AddAddress(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	chatID := update.Message.Chat.ID
	newAddress := update.Message.Text

	err := utils.SaveUserAddress(chatID, newAddress)
	if err != nil {
		return err
	}

	lang := utils.GetLanguage(update, config)
	msg := tgbotapi.NewMessage(chatID, utils.Translate(lang, "address_added"))
	_, err = bot.Send(msg)
	return err
}

func RemoveAddress(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	chatID := update.CallbackQuery.Message.Chat.ID
	addressID := update.CallbackQuery.Data

	err := utils.DeleteUserAddress(chatID, addressID)
	if err != nil {
		return err
	}

	lang := utils.GetLanguage(update, config)
	msg := tgbotapi.NewMessage(chatID, utils.Translate(lang, "address_removed"))
	_, err = bot.Send(msg)
	return err
}
