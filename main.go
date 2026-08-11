package main

import (
	"fmt"
)

func main() {
	if err := getToken(); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	Cities()
	var offset int = 0
	for {
		result, err := getUpdates(offset)
		if err != nil {
			fmt.Println(err)
		}
		for _, update := range result {
			chatID := update.Message.Chat.ID
			text := update.Message.Text
			if (text != "") && (chatID != 0) {
				text = handleMessage(chatID, text)
				err = sendMessage(chatID, text)
				if err != nil {
					fmt.Println(err)
				}

				offset = update.UpdateID + 1
			}
		}

	}
}
