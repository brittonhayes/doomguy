package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/brittonhayes/doomguy/internal"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	log "github.com/sirupsen/logrus"
	"time"
)

func (b *Bot) Setup() {
	b.Ctx.ChangeCommandInfo("Ping", "ping", "Check if the bot is responsive")
	b.Ctx.ChangeCommandInfo("Embed", "embed", "Create a custom rich embed")
	b.Ctx.ChangeCommandInfo("Say", "say", "Tell me what to say")
	b.Ctx.ChangeCommandInfo("Repeat", "repeat", "Repeat after you")
	b.Ctx.ChangeCommandInfo("Help", "help", "Get the available list of commands")
}

// Help prints the default help message.
func (b *Bot) Help(m *gateway.MessageCreateEvent) (*discord.Embed, error) {
	vars := map[string]interface{}{
		"Cmds": b.Ctx.Commands,
		"Subs": b.Ctx.Subcommands(),
	}

	output := b.Templates.ProcessFile("templates/help.tmpl", vars)
	// Make a new embed
	embed := discord.Embed{
		Title:       "⛑️ Help",
		Description: output,
		Footer:      &discord.EmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username)},
	}
	return &embed, nil
}

// Ping is a simple ping example, perhaps the most simple you could make it.
func (b *Bot) Ping(*gateway.MessageCreateEvent) (string, error) {
	return "Pong!", nil
}

// Say demonstrates how arguments.Flag could be used without the flag library.
func (b *Bot) Say(_ *gateway.MessageCreateEvent, f bot.ArgumentParts) (string, error) {
	if f.String() != "" {
		return f.String(), nil
	}
	return "", errors.New(internal.ErrorMissing)
}

// Repeat tells the bot to wait for the user's response, then repeat what they
// said.
func (b *Bot) Repeat(m *gateway.MessageCreateEvent) (string, error) {
	_, err := b.Ctx.SendMessage(m.ChannelID, "What do you want me to say?", nil)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// This might miss events that are sent immediately after. To make sure all
	// events are caught, ChanFor should be used.
	v := b.Ctx.WaitFor(ctx, func(v interface{}) bool {
		// Incoming event is a message create event:
		mg, ok := v.(*gateway.MessageCreateEvent)
		if !ok {
			return false
		}

		// Message is from the same author:
		return mg.Author.ID == m.Author.ID
	})

	if v == nil {
		e := errors.New(internal.ErrorTimeout)
		log.Error(e)
		return "", e
	}

	ev := v.(*gateway.MessageCreateEvent)
	return ev.Content, nil
}
