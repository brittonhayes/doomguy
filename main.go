package main

import (
	"github.com/brittonhayes/doomguy/pkg/base"
	"github.com/brittonhayes/doomguy/pkg/calc"
	"github.com/brittonhayes/doomguy/pkg/debug"
	"github.com/brittonhayes/doomguy/pkg/games"
	"github.com/brittonhayes/doomguy/pkg/guild"
	"github.com/brittonhayes/doomguy/pkg/reaction"
	"github.com/brittonhayes/doomguy/pkg/speaker"
	"github.com/brittonhayes/doomguy/pkg/template"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/gobuffalo/packr/v2"
	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = time.Stamp
	Formatter.ForceColors = true
	Formatter.FullTimestamp = true
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(Formatter)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var token = os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}

	// Load the templates from the
	// packr box
	box := packr.New("Templates", "./templates")
	tpl := template.NewTemplates(box)

	r := reaction.Reactions{}
	if err := configor.Load(&r, "reactions.yml"); err != nil {
		log.Fatal(err.Error())
	}

	commands := base.NewBot(tpl)
	wait, err := bot.Start(token, commands, func(ctx *bot.Context) error {
		// Setup pre-handlers
		r.React(ctx)
		// Setup user commands
		commands.Setup()
		ctx.HasPrefix = bot.NewPrefix("+")
		ctx.EditableCommands = true
		ctx.MustRegisterSubcommand(games.NewGamesClient(tpl))
		ctx.MustRegisterSubcommand(debug.NewDebug())
		ctx.MustRegisterSubcommand(calc.NewCalc())
		ctx.MustRegisterSubcommand(guild.NewGuild())
		ctx.MustRegisterSubcommand(speaker.NewSpeaker())
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Infof("Bot started")
	if err := wait(); err != nil {
		log.Fatalln("Gateway fatal error:", err)
	}
}
