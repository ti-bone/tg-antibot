/*
 * main.go
 * Copyright (c) ti-bone 2022.
 */

package main

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"net/http"
	"os"
	"tg-antibot/src/config"
	antibotHandlers "tg-antibot/src/handlers"
	"time"
)

func main() {
	// test mode
	var mode string

	if len(os.Args) != 1 {
		mode = os.Args[1]
	}

	testmode := false

	if mode == "--test" {
		config.LoadConfig("config/testconfig.json") // Load test config
		testmode = true
	} else {
		config.LoadConfig("config/config.json") // Load config
	}

	token := config.CurrentConfiguration.Bot.Token

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: 0,
			APIURL:  gotgbot.DefaultAPIURL,
		},
		UseTestEnvironment: testmode,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				fmt.Println("an error occurred while handling update:", err.Error())
				_, err2 := b.SendMessage(config.CurrentConfiguration.Bot.AdminID,
					fmt.Sprintf(
						"Error!\n\nMsg: %v\n\nFrom: %v\n\nChat: %v\n\nErr: %v",
						ctx.EffectiveMessage,
						ctx.EffectiveSender,
						ctx.EffectiveChat,
						err,
					), &gotgbot.SendMessageOpts{})

				if err2 != nil {
					fmt.Println(err2.Error())
				}

				return ext.DispatcherActionNoop
			},
			MaxRoutines: 1000,
		},
	})

	// Dispatcher
	dispatcher := updater.Dispatcher

	// Messages handler
	dispatcher.AddHandler(handlers.NewMessage(message.Text, antibotHandlers.Message)) // Text message

	// Joins handler
	dispatcher.AddHandler(handlers.NewMessage(message.NewChatMembers, antibotHandlers.Joins)) // Chat joins

	//Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 0,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 25,
			},
			AllowedUpdates: []string{
				"chat_member",
				"chat_join_request",
				"message",
			},
		},
	})

	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	fmt.Printf("@%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
