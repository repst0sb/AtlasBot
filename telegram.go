package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var TOKEN string

func getToken() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	TOKEN = os.Getenv("BOT_TOKEN")
	if TOKEN == "" {
		log.Fatal("Токен не найден в окружении")
	}
}

func getUpdates(offset int) ([]Update, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d&timeout=10", TOKEN, offset)
	r, err := http.Get(url)
	if err != nil {
		log.Println("Ошибка при getUpdates:", err)
		return nil, err
	}
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Println("Ошибка чтения body:", err)
		return nil, err
	}
	var resp Response
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		log.Println("Ошибка парсинга JSON:", err)
		return nil, err
	}
	if resp.Ok != true {
		return nil, fmt.Errorf("telegram API returned ok: false")
	}
	return resp.Result, nil
}

func sendMessage(chatID int64, text string, keyboard *InlineKeyboardMarkup) error {
	fmt.Printf("Получено сообщение от %d: %s\n", chatID, text)
	reqBody := &sendMessageReq{
		ChatID:       chatID,
		Text:         text,
		Reply_markup: keyboard,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("Ошибка парсинга JSON:", err)
		return err
	}
	postUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TOKEN)
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Println("Ошибка отправки:", err)
		return err
	} else {
		resp.Body.Close()
	}
	return nil
}

func answerCallbackQuery(callbackID string) error {
	reqBody := &answerCallbackQueryReq{CallbackQueryID: callbackID}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("Ошибка парсинга JSON:", err)
		return err
	}
	postUrl := fmt.Sprintf("https://api.telegram.org/bot%s/answerCallbackQuery", TOKEN)
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Println("Ошибка отправки:", err)
		return err
	} else {
		resp.Body.Close()
	}
	checkTelegramResponse(resp)
	return nil

}

func checkTelegramResponse(resp *http.Response) {
	fmt.Println(resp.StatusCode)
}
