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

type Response struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type CallbackQuery struct {
	ID      string  `json:"id"`
	From    User    `json:"from"`
	Message Message `json:"message"`
	Data    string  `json:"data"`
}

type User struct {
	ID         int64  `json:"id"`
	Is_bot     bool   `json:"is_bot"`
	First_name string `json:"first_name"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type sendMessageReq struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func getToken() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("ошибка загрузки .env: %w", err)
	}

	TOKEN = os.Getenv("BOT_TOKEN")
	if TOKEN == "" {
		return fmt.Errorf("токен не найден в окружении")
	}
	return err
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

func sendMessage(chatID int64, text string) error {
	fmt.Printf("Получено сообщение от %d: %s\n", chatID, text)
	reqBody := &sendMessageReq{
		ChatID: chatID,
		Text:   text,
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

func checkTelegramResponse(resp *http.Response) {
	fmt.Println(resp.StatusCode)
}
