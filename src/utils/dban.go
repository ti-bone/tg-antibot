/*
 * dban.go
 * Copyright (c) ti-bone 2022.
 */

package utils

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"html"
)

// DBan - Delete message and ban user/channel.
func DBan(b *gotgbot.Bot, ctx *ext.Context, reason string) error {
	// Check is message presents, if yes - delete message
	if ctx.Message != nil {
		// Delete message
		_, err := b.DeleteMessage(ctx.EffectiveChat.Id, ctx.Message.MessageId, &gotgbot.DeleteMessageOpts{})

		// Catch errors
		if err != nil {
			_, _ = b.SendMessage(ctx.EffectiveChat.Id, "Error while deleting message, please contact bot developer.", &gotgbot.SendMessageOpts{})
			return err
		}
	}

	// Check sender type
	if ctx.EffectiveSender.IsAnonymousChannel() {
		// Ban user from writing as channel
		_, err := b.BanChatSenderChat(ctx.EffectiveChat.Id, ctx.EffectiveSender.Id(), &gotgbot.BanChatSenderChatOpts{})

		// Catch errors
		if err != nil {
			_, _ = b.SendMessage(ctx.EffectiveChat.Id, "Error while banning channel, please contact bot developer.", &gotgbot.SendMessageOpts{})
			return err
		}
	} else if ctx.EffectiveSender.IsUser() {
		// Ban user in chat
		_, err := b.BanChatMember(ctx.EffectiveChat.Id, ctx.EffectiveSender.Id(), &gotgbot.BanChatMemberOpts{})

		// Catch errors
		if err != nil {
			_, _ = b.SendMessage(ctx.EffectiveChat.Id, "Error while banning user, please contact bot developer.", &gotgbot.SendMessageOpts{})
			return err
		}
	}

	// Send ban message
	_, err := b.SendMessage(ctx.EffectiveChat.Id,
		fmt.Sprintf(
			"<a href=\"tg://user?id=%d\">%s</a> [ID: <code>%d</code>] has been banned.\n<b>Reason:</b> %s",
			ctx.EffectiveSender.Id(),
			html.EscapeString(ctx.EffectiveSender.FirstName()),
			ctx.EffectiveSender.Id(),
			html.EscapeString(reason),
		),
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})

	// Catch errors
	if err != nil {
		_, _ = b.SendMessage(ctx.EffectiveChat.Id, "Error while sending message, please contact bot developer.", &gotgbot.SendMessageOpts{})
		return err
	}

	// Return no error
	return nil
}
