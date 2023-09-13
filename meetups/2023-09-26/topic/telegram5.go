keyboard := tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("button text 1", "some id to identify button 1"),
		tgbotapi.NewInlineKeyboardButtonData("button text 2", "some id to identify button 2"),
	),
)
msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Question, or Message")
msg.ReplyMarkup = keyboard
bot.Send(msg)
