package main

import (
	"time"
	"net/http"
	"fmt"
	"strconv"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

var WEBSITE_TO_MONITOR = os.Getenv("WEBSITE_TO_MONITOR")
var frequencyCheck string = os.Getenv("FREQUENCY_CHECK")
var TELEGRAM_TOKEN = os.Getenv("TELEGRAM_TOKEN")
var TELEGRAM_CHAT_ID, _ = strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)

func main() {
	freqCheck, _ := strconv.Atoi(frequencyCheck)
	for {
		rs, err := http.Get(WEBSITE_TO_MONITOR)
		if err != nil {
			sendTelegramMessage(err.Error())
		} else {
			if rs.StatusCode != http.StatusOK {
				sendTelegramMessage("statusCode: " + strconv.Itoa(rs.StatusCode) + " , statusMessage: " + rs.Status)
			}
		}
		time.Sleep(time.Duration(freqCheck) * time.Second)
	}
}

func sendTelegramMessage(message string) {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_TOKEN)
	if err != nil {
		fmt.Println(err)
	}

	msg := tgbotapi.NewMessage(TELEGRAM_CHAT_ID, message)

	bot.Send(msg)
}