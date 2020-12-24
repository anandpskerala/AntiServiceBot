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
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func antiExit(b *ext.Context) error {
	b.Bot.DeleteMessage(b.EffectiveChat.Id, b.EffectiveMessage.MessageId)
	return nil
}

func antiEntry(b *ext.Context) error {
	b.Bot.DeleteMessage(b.EffectiveChat.Id, b.EffectiveMessage.MessageId)
	return nil
}

func antiPin(b *ext.Context) error {
	b.Bot.DeleteMessage(b.EffectiveChat.Id, b.EffectiveMessage.MessageId)
	return nil
}

func antiMigrate(b *ext.Context) error {
	b.Bot.DeleteMessage(b.EffectiveChat.Id, b.EffectiveMessage.MessageId)
	return nil

}
func antiMigrateTo(b *ext.Context) error {
	b.Bot.DeleteMessage(b.EffectiveChat.Id, b.EffectiveMessage.MessageId)
	return nil
}

func antiMigrateFrom(b *ext.Context) error {
	b.Bot.DeleteMessage(b.EffectiveChat.Id, b.EffectiveMessage.MessageId)
	return nil
}
func LoadService(u ext.Updater) {
	u.Dispatcher.AddHandler(handlers.NewMessage(filters.NewChatMembers, antiEntry))
	u.Dispatcher.AddHandler(handlers.NewMessage(filters.LeftChatMembers, antiExit))
	u.Dispatcher.AddHandler(handlers.NewMessage(filters.Pin, antiPin))
	u.Dispatcher.AddHandler(handlers.NewMessage(filters.Migrate, antiMigrate))
	u.Dispatcher.AddHandler(handlers.NewMessage(filters.MigrateTo, antiMigrateTo))
	u.Dispatcher.AddHandler(handlers.NewMessage(filters.MigrateFrom, antiMigrateFrom))
}
