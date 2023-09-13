for update := range updates {
	var msg tgbotapi.MessageConfig
	if update.Message != nil {
		if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
			continue
		}
		switch update.Message.Text {
		case "/start":
			_, ok := cfg.AuthUsers[update.Message.From.ID]
			if !ok {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter token")
				msg.ReplyToMessageID = update.Message.MessageID

			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome back!")
				msg.ReplyToMessageID = update.Message.MessageID
			}
			bot.Send(msg)
		case "/login":
			_, ok := cfg.AuthUsers[update.Message.From.ID]
			if !ok {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter token:")
				msg.ReplyToMessageID = update.Message.MessageID

			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Already registered")
				msg.ReplyToMessageID = update.Message.MessageID
			}
			bot.Send(msg)
		case "/today":
			cmdToday(bot, cfg, update.Message)
		case "/yesterday":
			cmdYesterday(bot, cfg, update.Message)
		case "/date":
			_, ok := cfg.AuthUsers[update.Message.From.ID]
			if !ok {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hm, I do not know you")
				msg.ReplyToMessageID = update.Message.MessageID

			} else {
				// file := tgbotapi.FilePath("date.png")
				// bot.Send(file)
				// msg = tgbotapi.NewPhotoShare(update.Message.Chat.ID, "date.png")
				// pic := tgbotapi.NewChatPhoto(update.Message.Chat.ID, tgbotapi.FileURL("/home/zlatko/src/github.com/oktalz/pocs/huawei/date.jpg"))
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("09", "date:month:09"),
						tgbotapi.NewInlineKeyboardButtonData("08", "date:month:08"),
						tgbotapi.NewInlineKeyboardButtonData("07", "date:month:07"),
						tgbotapi.NewInlineKeyboardButtonData("06", "date:month:06"),
						tgbotapi.NewInlineKeyboardButtonData("05", "date:month:05"),
					),
				)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Pick a Month:")
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
				continue
			}
		case "/notification":
			_, ok := cfg.AuthUsers[update.Message.From.ID]
			if !ok {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hm, I do not know you")
				msg.ReplyToMessageID = update.Message.MessageID

			} else {
				pic := tgbotapi.NewPhoto(
					update.Message.From.ID,
					tgbotapi.FileBytes{
						Name:  "Not Implemented",
						Bytes: notImplemented,
					})
				bot.Send(pic)
				continue
			}
		default:
			found := false
			for _, user := range cfg.Users {
				if update.Message.Text == user.Token {
					user.ID = update.Message.From.ID
					cfg.AuthUsersFile[fmt.Sprintf("%d", user.ID)] = user
					cfg.AuthUsers[user.ID] = user
					var buf bytes.Buffer
					if err := toml.NewEncoder(&buf).Encode(cfg); err != nil {
						panic(err)
					}
					os.WriteFile("config.toml", buf.Bytes(), 0o644)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.FirstName+" "+update.Message.From.LastName+" excellent!!")
					msg.ReplyToMessageID = update.Message.MessageID
					found = true
					break
				}
			}
			if !found {
				pic := tgbotapi.NewPhoto(
					update.Message.From.ID,
					tgbotapi.FileBytes{
						Name:  "Not Implemented",
						Bytes: notImplemented,
					})
				bot.Send(pic)
				continue
			}
		}
	} else if update.CallbackQuery != nil {
		_, ok := cfg.AuthUsers[update.CallbackQuery.From.ID]
		if !ok {
			msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Hm, I do not know you")
			msg.ReplyToMessageID = update.Message.MessageID
			continue
		}
		// Respond to the callback query, telling Telegram to show the user
		// a message with the data received.
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if _, err := bot.Request(callback); err != nil {
			panic(err)
		}

		information := strings.Split(update.CallbackQuery.Data, ":")

		switch information[0] {
		case "/today":
			cmdToday(bot, cfg, update.CallbackQuery.Message)
		case "date":
			switch information[1] {
			case "month":
				monthStr := information[2]
				m, _ := strconv.Atoi(monthStr)
				keyboard := printMonthCalendar(2023, time.Month(m))
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "choose day")
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			case "day":
				// monthStr := information[3]
				// m, _ := strconv.Atoi(monthStr)
				date := information[4] + "-" + information[3] + "-" + information[2]
				layout := "2006-01-02"
				dt, err := time.Parse(layout, date)
				if err != nil {
					fmt.Println(err)
					return
				}
				mu.RLock()
				info, ok := database.Days[date]
				mu.RUnlock()
				if ok {
					img, err := os.ReadFile(path.Join("graph", date+".png"))
					if err != nil {
						pic := tgbotapi.NewPhoto(
							update.CallbackQuery.Message.Chat.ID,
							tgbotapi.FileBytes{
								Name:  "Not Implemented",
								Bytes: notImplemented,
							})
						bot.Send(pic)
						msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "missing image for "+date)
						bot.Send(msg)
					} else {
						pic := tgbotapi.NewPhoto(
							update.CallbackQuery.Message.Chat.ID,
							tgbotapi.FileBytes{
								Name:  date,
								Bytes: img,
							})
						bot.Send(pic)
					}
					//date:day:01:07:2023
					prevDate := dt.AddDate(0, 0, -1)
					nextDate := dt.AddDate(0, 0, 1)

					keyboard := tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData(prevDate.Format("02.01.2006"), "date:day:"+prevDate.Format("02:01:2006")),
							tgbotapi.NewInlineKeyboardButtonData(nextDate.Format("02.01.2006"), "date:day:"+nextDate.Format("02:01:2006")),
						),
					)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, dt.Format("02.01.2006")+": total power: "+info.String()+" kWh")
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				} else {
					pic := tgbotapi.NewPhoto(
						update.CallbackQuery.Message.Chat.ID,
						tgbotapi.FileBytes{
							Name:  "Not Implemented",
							Bytes: notImplemented,
						})
					bot.Send(pic)
				}
			default:
				pic := tgbotapi.NewPhoto(
					update.CallbackQuery.Message.Chat.ID,
					tgbotapi.FileBytes{
						Name:  "Not Implemented",
						Bytes: notImplemented,
					})
				bot.Send(pic)
			}
		default:
			pic := tgbotapi.NewPhoto(
				update.CallbackQuery.Message.Chat.ID,
				tgbotapi.FileBytes{
					Name:  "Not Implemented",
					Bytes: notImplemented,
				})
			bot.Send(pic)
		}
	}

}
