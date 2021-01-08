package games

import (
	"fmt"
	. "github.com/brittonhayes/doomguy/internal"
	"github.com/brittonhayes/doomguy/pkg/base"
	"github.com/brittonhayes/doomguy/pkg/template"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/dimuska139/rawg-sdk-go"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"math"
	"net/http"
	"time"
)

// Setup demonstrates the CanSetup interface. This function will never be parsed
// as a callback of any event.
func (g *Games) Setup(sub *bot.Subcommand) {
	// Set a custom command (e.g. "!go ..."):
	sub.Command = "games"
	// Set a custom description:
	sub.Description = "Fetch video game information"

	// Manually set the usage for each function.
	sub.ChangeCommandInfo("Get", "get", "Search for a game on RAWG")

	// Add aliases to ease up user experience
	// TODO fix aliases causing bot arguments to not remove
	sub.AddAliases("Get", "search", "find", "lookup")
}

// Get data about a specific game based on the user's query
func (g *Games) Get(m *gateway.MessageCreateEvent, query bot.ArgumentParts) (*discord.Embed, error) {
	if len(query) <= 0 {
		return nil, fmt.Errorf(template.Usage("games get subnautica"))
	}

	// Indicate that the bot is
	// working on a request
	Typing(m, g.Ctx)

	log.Infof("Searching game: '%s'", query.String())
	t := time.Now()

	client := g.Client
	filter := rawg.NewGamesFilter().
		SetSearch(query.String()).
		SetPage(1)

	data, _, err := client.GetGames(filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	details, err := client.GetGame(data[0].ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	embed := &base.CustomEmbed{
		Title:       fmt.Sprintf("%s 🕹️", details.Name),
		Description: g.Templates.GameDetails(details),
		Color:       "#00ffff",
		Image:       details.ImageBackground,
	}

	e, err := embed.NewEmbed()
	if err != nil {
		log.Error(err)
	}

	log.Infof("Search took %#v seconds", math.Round(time.Since(t).Seconds()))
	return e, nil
}

// NewGamesClient generates a new Games instance from the Games config
func NewGamesClient(t *template.Templates) *Games {
	g := new(Games).Config
	if err := configor.Load(&g, "games.yml"); err != nil {
		log.Fatal(err.Error())
	}

	config := rawg.Config{
		ApiKey:   g.ApiKey,
		Language: g.Language,
		Rps:      g.Rps,
	}

	return &Games{
		Templates: t,
		Client:    rawg.NewClient(http.DefaultClient, &config),
	}
}
