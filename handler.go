package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/go-zoox/fetch"
	"github.com/tidwall/gjson"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

func OnMessage(client *whatsmeow.Client, v *events.Message) {
	Info("Received a message from %s : %s", v.Info.Chat, v.Message.GetConversation())
	reply := ""
	words := strings.Split(v.Message.GetConversation(), " ")
	//command := regexp.MustCompile(`\.[a-zA-Z]+`).Split(text)
	command := words[0]
	query := words[1:]

	if command == "" {
		return
	}

	switch command {
	case ".info":
		reply += GetInfo()
	case ".menu":
		reply += GetInfo()
	case ".help":
		reply += GetInfo()
	case ".cuaca":
		reply += GetWeather(strings.Join(query, " "))
	case ".quotes":
		reply += GetQuotes()
	case ".ip":
		reply += IpLookup(query[0])
	default:
		reply = ""
	}

	reply = strings.TrimSpace(reply)

	if reply != "" {
		client.SendMessage(context.Background(), v.Info.Chat, &proto.Message{Conversation: &reply})
	}
}

func GetInfo() string {
	return `*- MENU -*
	
_.help .menu .info_
Bantuan untuk menggunakan bot ini
	
_.cuaca kota_
Lihat cuaca di kota tertentu

_.quotes_
Quotes random tiktok viral 2023

*- HENGKER -*
_.ip www.site.com_`
}

func GetWeather(location string) string {
	//format := url.QueryEscape("%l\n\n%C: %t%c\nAngin: %w\nBulan: %m\nTekanan: %P\nIndex UV: %u")
	//response, err := fetch.Get(fmt.Sprintf("https://wttr.in/%s?lang=id&format=%s", location, format))

	// https://api.openweathermap.org/data/2.5/forecast?q=Bandar%20Lampung&lang=id&units=metric&appid=20bd143d0d2383af677e195a47b89556
	response, err := fetch.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&lang=id&units=metric&appid=20bd143d0d2383af677e195a47b89556", location))
	json := string(response.Body)

	if err != nil || response.Status != 200 || !gjson.Valid(json) {
		Error("Cant get weather : %s", response.Error().Error())
		return "Gak bisa liat cuaca, coba lagi nanti"
	}

	lat := gjson.Get(json, "city.coord.lat").String()
	lon := gjson.Get(json, "city.coord.lon").String()

	s := location + "\n\n"
	limit := 5

	for i, v := range gjson.Get(json, "list").Array() {
		s += fmt.Sprintf("%s\n", v.Get("dt_txt").String())
		s += fmt.Sprintf("Suhu: %s°C\n", v.Get("main.temp").String())
		s += fmt.Sprintf("Cuaca: %s\n", v.Get("weather.0.description").String())
		s += "\n"

		if (i + 1) >= limit {
			break
		}
	}

	// give credit
	s += "_~ OpenWeatherMap_\n"
	s += fmt.Sprintf("http://www.google.com/maps/place/%s,%s", lat, lon)

	return s
}

func GetQuotes() string {
	// 	{
	// "_id": "VZAbqrJVc59C",
	// "content": "Friendship with oneself is all important because without it one cannot be friends with anybody else in the world.",
	// "author": "Eleanor Roosevelt",
	// "tags": [
	// "Famous Quotes"
	// ],
	// "authorSlug": "eleanor-roosevelt",
	// "length": 113,
	// "dateAdded": "2020-12-10",
	// "dateModified": "2023-04-14"
	// }

	response, err := fetch.Get(fmt.Sprintf("https://api.quotable.io/random"))
	json := string(response.Body)

	if err != nil || !gjson.Valid(json) {
		Error("Cant get quotes : %s", response.Error().Error())
		return "Gak bisa dapetin quote untukmu, coba lagi nanti"
	}

	quote := gjson.Get(json, "content").String()
	author := gjson.Get(json, "author").String()

	s := fmt.Sprintf("```%s```\n\n_~ %s_", quote, author)

	return s
}

func IpLookup(url string) string {
	Info("IP Lookup : %s", url)

	s := ""
	ips, _ := net.LookupIP(url)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			s += fmt.Sprintf("%s\n", ipv4)
		}
	}

	return s
}
