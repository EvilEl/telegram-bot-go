package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var botCommand = []tgbotapi.BotCommand{
	{
		Command:     "github",
		Description: "github",
	},
	{
		Command:     "guide",
		Description: "Справочник",
	},
}

func CreateCommands(bot *tgbotapi.BotAPI) error {
	setCommandsConfig := tgbotapi.NewSetMyCommands(botCommand...)
	_, err := bot.Request(setCommandsConfig)
	return err
}
