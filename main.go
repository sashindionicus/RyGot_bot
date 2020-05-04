package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const TgBotAPIKey = "1198171594:AAFCNWWGZhXdbdBu3meuf5udjLKNmVmQuIM"

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Главная 🙎‍♀"),
		tgbotapi.NewKeyboardButton("Запись ✍"),
		tgbotapi.NewKeyboardButton("Корзина 🛒"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Помощь 💁"),
	),
)

func main() {
	var (
		update    tgbotapi.Update
		updConfig tgbotapi.UpdateConfig
	)

	bot, err := tgbotapi.NewBotAPI(TgBotAPIKey)
	if err != nil {
		fmt.Printf("bot init error: %s\n", err)
		return
	}

	updConfig.Timeout = 60
	updConfig.Limit = 1
	updConfig.Offset = 0

	updChannel, err := bot.GetUpdatesChan(updConfig)
	if err != nil {
		fmt.Printf("update chan error: %s\n", err)
		return
	}

	for {
		update = <-updChannel
		if update.Message != nil {

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "test" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"test cmd")
					bot.Send(msgConfig)
				} else if cmdText == "menu" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
					msg.ReplyMarkup = mainMenu
					bot.Send(msg)
				}
			} else {
				fmt.Printf(
					"from: %s; chatID: %v; message: %s\n",
					update.Message.From.UserName,
					update.Message.Chat.ID,
					update.Message.Text)

				msgConfig := tgbotapi.NewMessage(
					update.Message.Chat.ID,
					update.Message.Text)
				bot.Send(msgConfig)
			}
		} else {
			fmt.Printf("not message: %+v\n", update)
		}
	}

	bot.StopReceivingUpdates()
}
