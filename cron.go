package main

import (
	"context"
	"fmt"
	"time"

	"github.com/goodsign/monday"
	"github.com/robfig/cron"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

var classesIn30Min = map[string]string{
	"Sen 06:30": "PANCASILA R-21 di GK1-306",
	"Sel 06:30": "TBFO RA di GK1-301",
	"Sel 09:00": "TBFO RB di GK1-302",
	"Rab 07:00": "OAK RC di GK1-303",
	"Rab 12:30": "MRV RB di GK1-302",
	"Kam 09:00": "MATDIS RB di GK1-302",
	"Jum 09:00": "MATDIS RB di GK1-302",
}

func StartCron() {
	//loc := time.FixedZone("UTC+7", 7*60*60)
	c := cron.New()
	c.AddFunc("0 * * * * *", CheckTask)
	c.Start()
	Info("Cron started")
}

func CheckTask() {
	check := GetDateToCheck()
	class, ok := classesIn30Min[check]

	Info("Checking schedule for %s", check)
	Info("Found : %b", ok)

	if ok {
		text := fmt.Sprintf("Reminder : kelas %s dalam 30 menit", class)
		client.SendMessage(context.Background(), types.NewJID("120363028653412122", types.GroupServer), &proto.Message{
			Conversation: &text,
		})
		return
	}
}

func GetDateToCheck() string {
	currentTime := time.Now().UTC().Add(7 * time.Hour)
	shortDate := monday.Format(currentTime, "Mon", monday.LocaleIdID)
	check := fmt.Sprintf("%s %d:%d", shortDate, currentTime.Hour(), currentTime.Minute())
	return check
}
