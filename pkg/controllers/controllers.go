package controllers

import (
	"log"

	"github.com/fi9ish/binanceInfoBot/pkg/binanceRequests"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update, command string, args string) {
	if command == "GetCoins" {
		binanceRequests.GetCoins()
	}
}

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, message string) {
	response := message
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
