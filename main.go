package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
	"fmt"
)

var (
	windDir = map[string]string{
		"nw": "северо-западный",
		"n":  "северный",
		"ne": "северо-восточный",
		"e":  "восточный",
		"se": "юго-восточный",
		"s":  "южный",
		"sw": "юго-западный",
		"w":  "западный",
		"с":  "штиль",
	}

	conditions = map[string]string{
		"clear":                  "ясно",
		"partly-cloudy":          "малооблачно",
		"cloudy":                 "облачно с прояснениями",
		"overcast":               "пасмурно",
		"drizzle":                "морось",
		"light-rain":             "небольшой дождь",
		"rain":                   "дождь",
		"moderate-rain":          "умеренно сильный дождь",
		"heavy-rain":             "сильный дождь",
		"continuous-heavy-rain":  "длительный сильный дождь",
		"showers":                "ливень",
		"wet-snow":               "дождь со снегом",
		"light-snow":             "небольшой снег",
		"snow":                   "снег",
		"snow-showers":           "снегопад",
		"hail":                   "град",
		"thunderstorm":           "гроза",
		"thunderstorm-with-rain": "дождь с грозой",
		"thunderstorm-with-hail": "гроза с градом",
	}
)

const (
	webHook = "https://gz-weather-bot.herokuapp.com"
)

func main() {
	port := os.Getenv("PORT")
	publicURL := os.Getenv("PUBLIC_URL")
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("токен бота должен быть указан в качестве значения переменной окружения BOT_TOKEN")
	}

	yandexToken := os.Getenv("YANDEX_TOKEN")
	if yandexToken == "" {
		log.Fatal("токен бота должен быть указан в качестве значения переменной окружения YANDEX_TOKEN")
	}
	
	pref := tb.Settings{
		Token:  botToken,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	api := NewAPI(yandexToken)

	b.Handle(tb.OnText, func(m *tb.Message) {
		city := m.Text
		resp, _ := api.Request(city)
		mess := fmt.Sprintf("Температура: %d Ощущается: %d\nВетер: %0.2fмс %s\n%s", resp.Fact.Temp, resp.Fact.Feels, resp.Fact.WindSpeed, windDir[resp.Fact.WindDir], conditions[resp.Fact.Condition])
		b.Send(m.Sender, mess)
	})

	b.Start()
}
