package games

import (
	"github.com/brittonhayes/doomguy/pkg/template"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/dimuska139/rawg-sdk-go"
)

var _ Gamer = (*Games)(nil)

// Gamer contains the methods available
// to the Games type
type Gamer interface {
	Get(m *gateway.MessageCreateEvent, query bot.ArgumentParts) (*discord.Embed, error)
}

// Games contains the fields used to query the RAWG REST api
// and return information about video games
type Games struct {
	// Context must not be embedded.
	Ctx       *bot.Context
	Client    *rawg.Client
	Templates *template.Templates
	Config    struct {
		ApiKey   string   `yaml:"api_key" default:""`
		Language string   `yaml:"language" default:"en"`
		Rps      int      `yaml:"rate" default:"5"`
		Aliases  []string `yaml:"aliases"`
	}
}
