package bot

import (
	"fmt"
	"log"
	"myapp/handlers/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateBot() {

	config := config.New()
	fmt.Println("CreateBot:config.TELEGRAM_APITOKEN", config.TELEGRAM_APITOKEN)
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_APITOKEN)

	if err != nil {
		panic(err)
	}
	err = CreateCommands(bot)

	if err != nil {
		log.Fatalf("Ошибка установки команд: %v", err)
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "markdown"
		// Extract the command from the Message.
		switch update.Message.Command() {
		case "github":
			msg.Text = GetGithubLink()
		case "guide":
			msg.Text = GetGuideText()
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	bot.Debug = true
}
