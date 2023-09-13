if update.CallbackQuery != nil {
	_, ok := cfg.AuthUsers[update.CallbackQuery.From.ID]
	...
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		continue
	}

	switch update.CallbackQuery.Data {
	case ...:
		...
	}
}
