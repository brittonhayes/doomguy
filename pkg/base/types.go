package base

import (
	"github.com/brittonhayes/doomguy/pkg/template"
	"github.com/diamondburned/arikawa/v2/bot"
)

// Bot is the primary struct for all bot commands
type Bot struct {
	// Context must not be embedded.
	Ctx       *bot.Context
	Templates *template.Templates
}

func NewBot(tpl *template.Templates) *Bot {
	return &Bot{
		Templates: tpl,
	}
}
