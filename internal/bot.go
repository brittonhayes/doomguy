package internal

import (
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	log "github.com/sirupsen/logrus"
)

// Indicate that the bot is working on the request
func Typing(m *gateway.MessageCreateEvent, ctx *bot.Context) {
	// Indicate bot is typing
	if err := ctx.Typing(m.ChannelID); err != nil {
		log.Error(err.Error())
	}
}
