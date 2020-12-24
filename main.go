/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/antiservicebot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/github.com/Anandpskerala/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package main

import (
	"fmt"
	"os"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/anandpskerala/antiservicebot/service"
)

func main() {

	token := os.Getenv("TOKEN")
	b, err := gotgbot.NewBot(token)
	if err != nil {
		log.Fatalf("New bot creation failed", err.Error())
		return
	}

	u := ext.NewUpdater(b, nil)

	service.LoadService(u)
	u.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	err = u.StartPolling(b, &ext.PollingOpts{Clean: true})
	if err != nil {
		log.Fatalf("Polling failed", err.Error())
		return
	}
	bot := u.Bot.User.FirstName
	fmt.Printf("Successfully logged as %s", bot)

	u.Idle()

}

func start(b *ext.Context) error {

	_, err := b.Bot.SendMessage(b.EffectiveChat.Id, "Hi Buddy, I am <b>AntiServiceMessageBot</b>\n\nI am a bot which can delete service message like when a user <u>enters</u> or <u>exists</u> a group\n\nI am a fully written in <b>Go</b>\n\n<i>Note</i> :  You should promote me as an admin and give atleast two admins rights shown below for getting my full service:\n * Right to Deleted Messages\n * Right to add admins\n\n<b>Support Group</b> : <b><a href='https://t.me/Keralasbots'>KeralaBots</a></b>", &gotgbot.SendMessageOpts{ReplyToMessageId: b.EffectiveMessage.MessageId, ParseMode: "HTML"})
	if err != nil {
		log.Printf("Error in sending", err.Error())
	}
	return err
}
