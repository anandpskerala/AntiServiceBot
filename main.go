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

	//"github.com/anandpskerala/antiservicebot/config"
	"github.com/anandpskerala/antiservicebot/service"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/parsemode"

	//"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := zap.NewProductionEncoderConfig()

	cfg.EncodeLevel = zapcore.CapitalLevelEncoder

	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.InfoLevel))

	defer logger.Sync()

	l := logger.Sugar()

	token := os.Getenv("TOKEN")
	u, err := gotgbot.NewUpdater(logger, token)
	if err != nil {
		l.Fatalw("Updater failed starting", zap.Error(err))
		return
	}

	service.LoadService(u)

	u.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	err = u.StartPolling()
	if err != nil {
		l.Fatalw("Polling failed", zap.Error(err))
		return
	}
	bot := u.Bot.FirstName
	fmt.Printf("Successfully logged as %s", bot)

	u.Idle()

}

func start(b ext.Bot, u *gotgbot.Update) error {

	msg := b.NewSendableMessage(u.EffectiveChat.Id, "Hi Buddy, I am <b>AntiServiceMessageBot</b>\n\nI am a bot which can delete service message like when a user <u>enters</u> or <u>exists</u> a group\n\nI am a fully written in <b>Go</b>\n\n<i>Note</i> :  You should promote me as an admin and give atleast two admins rights shown below for getting my full service:\n * Right to Deleted Messages\n * Right to add admins\n\n<b>Support Group</b> : <b><a href='https://t.me/Keralasbots'>KeralaBots</a></b>")
	msg.ReplyToMessageId = u.EffectiveMessage.MessageId
	msg.ParseMode = parsemode.Html
	_, err := msg.Send()
	if err != nil {
		b.Logger.Warnw("Error in sending", zap.Error(err))
	}
	return err
}
