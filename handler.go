package main

import (
	"fmt"
	"strings"
	"time"
)

var userStates = make(map[int64]*RouteInfo)
var cities = make(map[string]string)

type RouteInfo struct {
	From     string
	To       string
	Data     string
	TimeFrom time.Time
	TimeTo   time.Time
}

func handleMessage(chatID int64, text string) (answer string) {
	if text == "/start" {
		answer = "Выберите город откуда едите, куда едите, дату и время. Например Минск Микашевичи 2026-08-07 12:00-15:00 "
		return answer
	}
	var userRoute RouteInfo
	err := GetRouteInfo(text, &userRoute)
	if err != nil {
		fmt.Println(err)
		answer = "ошибка"
		return answer
	}
	ride, err := AtlasRequest(userRoute.From, userRoute.To, userRoute.Data)
	if err != nil {
		fmt.Println("что-то с атласом")
		answer = "ошибка атлас"
	}
	answer = FormMessage(ride)
	userStates[chatID] = &userRoute
	return answer
}

func FormMessage(a *AtlasResponse) string {
	var s string
	for _, v := range a.Rides {
		s += fmt.Sprintf("%s -> %s \n	Мест: %d\n	Цена: %d\n", v.Arrival, v.Departure, v.FreeSeats, v.Price)
	}
	return s
}

func GetRouteInfo(text string, b *RouteInfo) error {
	words := strings.Fields(text)
	if len(words) != 4 {
		return fmt.Errorf("Неверный запрос")
	}
	city := findCities(words[0])
	if city == "" {
		return fmt.Errorf("такого города отправления нет")
	}
	b.From = city
	city = findCities(words[1])
	if city == "" {
		return fmt.Errorf("такого города прибытия нет")
	}
	b.To = city
	_, err := time.Parse("2006-01-02", words[2])
	if err != nil {
		return fmt.Errorf("неверная дата")
	}
	b.Data = words[2]
	times := strings.Split(words[3], "-")
	b.TimeFrom, err = time.Parse("15:04", times[0])
	if err != nil {
		return fmt.Errorf("неверное время")
	}
	b.TimeFrom = b.TimeFrom.In(time.Local)
	b.TimeTo, err = time.Parse("15:04", times[1])
	if err != nil {
		return fmt.Errorf("неверное время")
	}
	return nil
}

func findCities(city string) string {
	for i, v := range cities {
		if city == i {
			return v
		}
	}
	return ""
}

func Cities() {
	cities["Микашевичи"] = "c625144"
	cities["Минск"] = "c625228"
}
