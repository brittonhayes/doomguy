package guild

import "github.com/diamondburned/arikawa/v2/bot"

// Flag for administrators only.
type Guild struct {
	Ctx *bot.Context
}

// NewGuild generates a new instance of Guild
func NewGuild() *Guild {
	return &Guild{}
}
