package telegram

import (
	"fmt"
	"shop-bot/internal/delivery/telegram/buttons"
	"shop-bot/internal/delivery/telegram/utils"
	"shop-bot/internal/usecase/interfaces"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewHandleCart(cartUseCase interfaces.CartUseCase) func(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
	return func(update tgbotapi.Update, bot *tgbotapi.BotAPI, config *utils.TelegramConfig) error {
		lang := utils.GetLanguage(update, config)

		userID := update.Message.Chat.ID
		cart, err := cartUseCase.GetCartByUserID(userID)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "error_fetching_cart"))
			_, _ = bot.Send(msg)
			return err
		}

		if len(cart.Items) == 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, utils.Translate(lang, "cart_empty"))
			_, _ = bot.Send(msg)
			return nil
		}

		response := fmt.Sprintf("%s\n\n", utils.Translate(lang, "cart_content"))
		for _, item := range cart.Items {
			response += fmt.Sprintf(utils.Translate(lang, "cart_line_item"), item.CartID, item.Quantity, item.Price*float64(item.Quantity))
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyMarkup = buttons.CartMenuKeyboard(lang)
		_, err = bot.Send(msg)
		return err
	}
}
