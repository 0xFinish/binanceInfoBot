package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/fi9ish/binanceInfoBot/pkg/config"
	"github.com/fi9ish/binanceInfoBot/pkg/controllers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	fmt.Println("how are we doing fam?")
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		fmt.Println("Not found error")
		log.Fatal(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}
			if update.Message.IsCommand() {
				command := update.Message.Command()
				args := update.Message.CommandArguments()

				controllers.HandleCommand(bot, update, command, args)
			} else {
				message := update.Message.Text

				// Do something with the message
				controllers.HandleMessage(bot, update, message)
			}
		}
	}()

	select {}
}
