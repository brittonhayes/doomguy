package reaction

import (
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/utils/handler"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	SuccessMessage = "Successfully reacted to: "
)

// React to any of the listed phrases with an emote
func (r *Reactions) React(ctx *bot.Context) {
	ctx.PreHandler = handler.New()
	ctx.PreHandler.Synchronous = true
	ctx.PreHandler.AddHandler(func(c *gateway.MessageCreateEvent) {
		for _, trigger := range r.Config {
			for _, t := range trigger.Phrases {
				if strings.Contains(c.Message.Content, t) {
					err := ctx.React(c.ChannelID, c.Message.ID, discord.APIEmoji(trigger.Emoji))
					if err != nil {
						log.Error(err.Error())
					}
					log.Info(SuccessMessage, t)
				}
			}
		}
	})
	ctx.Gateway.AddIntents(gateway.IntentGuildMessages)
	ctx.Gateway.AddIntents(gateway.IntentGuildMessageReactions)
	if err := ctx.Open(); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer ctx.Close()
}
