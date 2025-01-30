package telegram

import (
	"fmt"
	"shop-bot/internal/delivery/telegram/buttons"
	"shop-bot/internal/delivery/telegram/utils"
	"shop-bot/internal/usecase/interfaces"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleSettings(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig, userUseCase interfaces.UserUseCase) error {
	lang := utils.GetLanguage(update, config)

	userID := update.Message.Chat.ID
	user, err := userUseCase.GetUserByID(userID) // Метод GetProfile должен быть переименован на GetUserProfile
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "error_fetching_user"))
		_, _ = bot.Send(msg)
		return err
	}

	response := fmt.Sprintf(
		"%s\n\n"+utils.Translate(lang, "profile_name")+"\n"+utils.Translate(lang, "profile_phone"),
		utils.Translate(lang, "profile_section"),
		user.Username,
		user.Email,
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ReplyMarkup = buttons.SettingsMenuKeyboard(lang)
	_, err = bot.Send(msg)
	return err
}
