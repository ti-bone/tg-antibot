/*
 * joins.go
 * Copyright (c) ti-bone 2022.
 */

package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"tg-antibot/src/filters"
	"tg-antibot/src/utils"
)

// Joins - Handler for new members
func Joins(b *gotgbot.Bot, ctx *ext.Context) error {
	// Fetch info about sender
	chatMember, err := b.GetChatMember(ctx.EffectiveChat.Id, ctx.EffectiveSender.Id(), &gotgbot.GetChatMemberOpts{})

	// Catch errors
	if err != nil {
		_, err2 := ctx.EffectiveMessage.Reply(b, "Error while fetching info about sender, please contact bot developer.", &gotgbot.SendMessageOpts{})
		if err2 != nil {
			return err2
		}
		return err
	}

	// Ignore chat admins and banned users
	if chatMember.GetStatus() == "administrator" || chatMember.GetStatus() == "creator" || chatMember.GetStatus() == "kicked" {
		return nil
	}

	// Fetch current filters
	curFilters := filters.GetFilters()

	// Name filtering
	if utils.Compare(ctx.EffectiveSender.Name(), curFilters.Names) {
		return utils.DBan(b, ctx, "User's name contains spam-characters")
	}

	// Return no error
	return nil
}
