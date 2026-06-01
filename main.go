package main

import (
	"fmt"
)

func main() {
	getToken()
	var offset int = 0
	for {
		result, err := getUpdates(offset)
		if err != nil {
			fmt.Println(err)
		}
		for _, update := range result {
			var keyboard *InlineKeyboardMarkup
			chatID := update.Message.Chat.ID
			text := update.Message.Text
			if (text != "") && (chatID != 0) {
				text, keyboard = handleMessage(chatID, text)
			} else if (update.Callback_query.From.ID != 0) && (update.Callback_query.Data != "") {
				text, keyboard = handleCallback(update.Callback_query)
				chatID = update.Callback_query.Message.Chat.ID
				err = answerCallbackQuery(update.Callback_query.ID)
				if err != nil {
					fmt.Println(err)
				}
			}
			err = sendMessage(chatID, text, keyboard)
			if err != nil {
				fmt.Println(err)
			}

			offset = update.UpdateID + 1
		}

	}
}
