/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/antiservicebot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/github.com/Anandpskerala/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package service

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
)

func antiExit(b ext.Bot, u *gotgbot.Update) error {
	b.DeleteMessage(u.EffectiveChat.Id, u.EffectiveMessage.MessageId)
	return nil
}

func antiEntry(b ext.Bot, u *gotgbot.Update) error {
	b.DeleteMessage(u.EffectiveChat.Id, u.EffectiveMessage.MessageId)
	return nil
}

func antiPin(b ext.Bot, u *gotgbot.Update) error {
	b.DeleteMessage(u.EffectiveChat.Id, u.EffectiveMessage.MessageId)
	return nil
}

func antiMigrate(b ext.Bot, u *gotgbot.Update) error {
	b.DeleteMessage(u.EffectiveChat.Id, u.EffectiveMessage.MessageId)
	return nil

}
func antiMigrateTo(b ext.Bot, u *gotgbot.Update) error {
	b.DeleteMessage(u.EffectiveChat.Id, u.EffectiveMessage.MessageId)
	return nil
}

func antiMigrateFrom(b ext.Bot, u *gotgbot.Update) error {
	b.DeleteMessage(u.EffectiveChat.Id, u.EffectiveMessage.MessageId)
	return nil
}
func LoadService(u *gotgbot.Updater) {
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.NewChatMembers(), antiEntry))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.LeftChatMembers(), antiExit))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Pin, antiPin))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Migrate(), antiMigrate))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.MigrateFrom(), antiMigrateTo))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.MigrateTo(), antiMigrateFrom))
}
