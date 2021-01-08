package reaction

import "github.com/diamondburned/arikawa/v2/bot"

var _ Reactor = (*Reactions)(nil)

// Reactor contains the methods available
// to the Reactions type
type Reactor interface {
	React(ctx *bot.Context)
}

// Reactions
type Reactions struct {
	Ctx    *bot.Context
	Config []struct {
		Emoji   string   `yaml:"emoji"`
		Phrases []string `yaml:"phrases"`
	}
}
