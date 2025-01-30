package utils

func Translate(lang, key string) string {
	translations := map[string]map[string]string{
		"en": {
			"welcome_message":       "Welcome to the bot!",
			"help_message":          "Here is how I can help you:",
			"order_button":          "Order",
			"cart_button":           "Cart",
			"contacts_button":       "Contacts",
			"settings_button":       "Settings",
			"back_button":           "Back",
			"cart_empty":            "Your cart is empty.",
			"cart_content":          "Here is your cart content:",
			"cart_clear_all":        "Clear cart",
			"cart_checkout":         "Checkout",
			"cart_line_item":        "%s x %d - %0.2f USD\n",
			"cart_cleared":          "Cart cleared successfully!",
			"error_fetching_cart":   "Failed to fetch your cart. Please try again later.",
			"settings_language":     "Select your language:",
			"profile_section":       "Profile information:",
			"profile_name":          "Name: %s",
			"profile_phone":         "Phone: %s",
			"profile_update_prompt": "Send me new data to update.",
		},
		"ru": {
			"welcome_message":       "Добро пожаловать в бота!",
			"help_message":          "Вот чем я могу вам помочь:",
			"order_button":          "Заказ",
			"cart_button":           "Корзина",
			"contacts_button":       "Контакты",
			"settings_button":       "Настройки",
			"back_button":           "Назад",
			"cart_empty":            "Ваша корзина пуста.",
			"cart_content":          "Содержимое вашей корзины:",
			"cart_clear_all":        "Очистить корзину",
			"cart_checkout":         "Оформить заказ",
			"cart_line_item":        "%s x %d - %0.2f рублей\n",
			"cart_cleared":          "Корзина успешно очищена!",
			"error_fetching_cart":   "Не удалось загрузить вашу корзину. Попробуйте позже.",
			"settings_language":     "Выберите язык:",
			"profile_section":       "Информация профиля:",
			"profile_name":          "Имя: %s",
			"profile_phone":         "Телефон: %s",
			"profile_update_prompt": "Отправьте новые данные для обновления.",
		},
	}

	if val, ok := translations[lang][key]; ok {
		return val
	}
	if val, ok := translations["en"][key]; ok {
		return val
	}
	return key
}

func GetLanguage(update any, config *TelegramConfig) string {
	if config.DefaultLanguage != "" {
		return config.DefaultLanguage
	}
	return "en"
}
