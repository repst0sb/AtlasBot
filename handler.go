package main

var userStates = make(map[int64]*UserState)

const RouteInMikashi = "Минск -> Микашевичи"
const RouteInMinsk = "Микашевичи -> Минск"
const TimeMorning = "5:00-9:00"
const TimeAfternoon = "9:00-14:00"
const TimeEvening = "14:00-18:00"
const TimeNight = "18:00-23:00"

func handleMessage(chatID int64, text string) (answer string, keyboard *InlineKeyboardMarkup) {
	if text == "/start" {
		answer = "Выберите откуда куда поедем?"
		buttons := [][]InlineKeyboardButton{
			{
				{Text: RouteInMinsk, Callback_data: RouteInMinsk},
			}, {
				{Text: RouteInMikashi, Callback_data: RouteInMikashi},
			},
		}
		keyboard = &InlineKeyboardMarkup{Inline_keyboard: buttons}
		return answer, keyboard
	}
	answer = "Я не понимаю. Напишите /start для начала работы со мной."
	return answer, nil
}

func handleCallback(query CallbackQuery) (answer string, keyboard *InlineKeyboardMarkup) {
	user := userStates[query.Message.Chat.ID]
	if (user != nil) && (user.Route != "") && (user.TimeRange != "") {
		answer = "Вы уже запустили поиск, ожидайте!"
		return answer, nil
	}
	if (query.Data == RouteInMinsk) || (query.Data == RouteInMikashi) {
		userStates[query.Message.Chat.ID] = &UserState{
			Route: query.Data,
		}
		answer = "Выберите время для маршрута: " + query.Data
		buttons := [][]InlineKeyboardButton{
			{
				{Text: TimeMorning, Callback_data: TimeMorning},
			}, {
				{Text: TimeAfternoon, Callback_data: TimeAfternoon},
			}, {
				{Text: TimeEvening, Callback_data: TimeEvening},
			}, {
				{Text: TimeNight, Callback_data: TimeNight},
			},
		}
		keyboard = &InlineKeyboardMarkup{Inline_keyboard: buttons}
		return answer, keyboard
	}
	if user == nil {
		answer = "Ошибка! Повторите еще раз."
		return answer, nil
	}
	userStates[query.Message.Chat.ID].TimeRange = query.Data
	answer = "Начинаю поиск маршруток. Маршрут: " + user.Route + ". Время " + user.TimeRange
	return answer, nil
}
