package guild

import (
	"fmt"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/bot/extras/middlewares"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// Setup demonstrates the CanSetup interface. This function will never be parsed
// as a callback of any event.
func (guild *Guild) Setup(sub *bot.Subcommand) {
	// Set a custom command (e.g. "!go ..."):
	sub.Command = "guild"
	// Set a custom description:
	sub.Description = "Interact with guild details"

	// Restrict commands to guild
	sub.AddMiddleware("GuildInfo", middlewares.GuildOnly(guild.Ctx))

	// Manually set the usage for each function.
	sub.ChangeCommandInfo("GuildInfo", "info", "Get the guild's information")
}

// GuildInfo demonstrates the GuildOnly middleware done in (*Bot).Setup().
func (guild *Guild) GuildInfo(m *gateway.MessageCreateEvent) (*discord.Embed, error) {

	color := "#8a0303"
	g, err := guild.Ctx.GuildWithCount(m.GuildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild: %v", err)
	}

	colorHex, err := strconv.ParseInt((color)[1:], 16, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse color, %s", color)
	}

	// Make a new embed
	embed := discord.Embed{
		Title:       g.Name,
		Image:       &discord.EmbedImage{URL: g.IconURL()},
		Description: fmt.Sprintf("ID: %v\nMembers: %v\n", g.ID, g.ApproximateMembers),
		Color:       discord.Color(colorHex),
	}
	log.Info("Discord guild info requested")

	return &embed, nil
}
