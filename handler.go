package main

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

func OnMessage(client *whatsmeow.Client, v *events.Message) {
	Info("Received a message from %s : %s", v.Info.Chat, v.Message.Conversation)
	reply := ""
	text := v.Message.GetConversation()

	switch text {
	case ".info":
		reply += "*INFO*\n"
		reply += ".halo\n"
		reply += ".cuaca\n"
		reply += ".quotes\n"
	case ".halo":
		reply += fmt.Sprintf("Halo juga, %s", v.Info.PushName)
	case ".cuaca":
		reply += "Udahlah pasti panas malah pake nanya"
	case ".quotes":
		reply += "Gak perlu quotes, sana cuci piringnya"
	default:
		reply = ""
	}

	if reply != "" {
		client.SendMessage(context.Background(), v.Info.Chat, &proto.Message{Conversation: &reply})
	}
}
