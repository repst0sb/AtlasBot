package main

type Response struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateID       int           `json:"update_id"`
	Message        Message       `json:"message"`
	Callback_query CallbackQuery `json:"callback_query"`
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

type answerCallbackQueryReq struct {
	CallbackQueryID string `json:"callback_query_id"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type sendMessageReq struct {
	ChatID       int64                 `json:"chat_id"`
	Text         string                `json:"text"`
	Reply_markup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineKeyboardMarkup struct {
	Inline_keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text          string `json:"text"`
	Callback_data string `json:"callback_data"`
}
type UserState struct {
	Route     string
	TimeRange string
}

type AtlasResponse struct {
	Rides []Ride `json:"rides"`
}

type Ride struct {
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	FreeSeats string `json:"freeSeats"`
	Price     string `json:"price"`
}
