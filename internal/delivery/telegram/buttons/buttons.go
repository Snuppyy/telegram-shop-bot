package buttons

import (
	"shop-bot/internal/delivery/telegram/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MainMenuKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "order_button")+" 🛒"),
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "cart_button")+" 🛍️"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "contacts_button")+" 📞"),
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "settings_button")+" ⚙️"),
		),
	)
}

func SettingsMenuKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "change_language")),
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "edit_profile")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "manage_addresses")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "back_button")+" 🔙"),
		),
	)
}

func OrderMenuKeyboard(lang string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(utils.Translate(lang, "categories_button")+" 📂", "categories"),
			tgbotapi.NewInlineKeyboardButtonData(utils.Translate(lang, "search_button")+" 🔍", "search"),
		),
	)
}

func AddressMenuKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "add_address")),
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "remove_address")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "address_list")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "back_button")+" 🔙"),
		),
	)
}

func LanguageKeyboard(languages []string) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, lang := range languages {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lang, lang),
		)
		rows = append(rows, row)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func CartMenuKeyboard(lang string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(utils.Translate(lang, "clear_cart")+" 🗑️", "clear_cart"),
			tgbotapi.NewInlineKeyboardButtonData(utils.Translate(lang, "checkout")+" 💳", "checkout"),
		),
	)
}

func BackToSettingsKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(utils.Translate(lang, "back_button") + " 🔙"),
		),
	)
}
