for update := range updates {
	if update.Message != nil {
		_, ok := cfg.AuthUsers[update.Message.From.ID]
		...
		
	} else if update.CallbackQuery != nil {
		_, ok := cfg.AuthUsers[update.CallbackQuery.From.ID]
		...
	}
	...
}
