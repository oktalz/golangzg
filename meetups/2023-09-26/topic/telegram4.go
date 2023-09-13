pic := tgbotapi.NewPhoto(
	update.CallbackQuery.Message.Chat.ID,
	tgbotapi.FileBytes{
		Name:  "Image Name",
		Bytes: img,
	})
bot.Send(pic)
