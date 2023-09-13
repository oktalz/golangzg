
switch update.Message.Text {
case "/start":
	...
case "/login":
	...
case "/today":
	cmdToday(bot, cfg, update.Message)
case "/yesterday":
	cmdYesterday(bot, cfg, update.Message)
case "/date":
	...
