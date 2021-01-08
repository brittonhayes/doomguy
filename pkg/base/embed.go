package base

import (
	"errors"
	"fmt"
	"github.com/diamondburned/arikawa/v2/bot/extras/arguments"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"strconv"
	"strings"
)

type CustomEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Author      string `json:"author"`
	Footer      string `json:"footer"`
	Color       string `json:"color"`
}

// Embed is a simple embed creator. Its purpose is to demonstrate the usage of
// the ParseContent interface, as well as using the stdlib flag package.
func (b *Bot) Embed(_ *gateway.MessageCreateEvent, f arguments.Flag) (*discord.Embed, error) {
	fs := arguments.NewFlagSet()

	var (
		title  = fs.String("title", "", "Title of the embed")
		author = fs.String("author", "", "Author of the embed")
		footer = fs.String("footer", "", "Footer of the embed")
		color  = fs.String("color", "#FFFFFF", "Color in hex format #hhhhhh")
	)

	if err := f.With(fs.FlagSet); err != nil {
		return nil, err
	}

	if len(fs.Args()) < 1 {
		return nil, fmt.Errorf("usage: embed [flags] content...\n" + fs.Usage())
	}

	// Check if the color string is valid.
	if !strings.HasPrefix(*color, "#") || len(*color) != 7 {
		return nil, errors.New("invalid color, format must be #hhhhhh")
	}

	// Parse the color into decimal numbers.
	colorHex, err := strconv.ParseInt((*color)[1:], 16, 64)
	if err != nil {
		return nil, err
	}

	// Make a new embed
	embed := discord.Embed{
		Title:       *title,
		Description: strings.Join(fs.Args(), " "),
		Color:       discord.Color(colorHex),
	}

	if *author != "" {
		embed.Author = &discord.EmbedAuthor{
			Name: *author,
		}
	}
	if *footer != "" {
		embed.Footer = &discord.EmbedFooter{
			Text: *footer,
		}
	}

	return &embed, err
}

// NewEmbed generates an embed from a CustomEmbed struct for internal methods
func (c *CustomEmbed) NewEmbed() (*discord.Embed, error) {
	// Check if the color string is valid.
	if !strings.HasPrefix(c.Color, "#") || len(c.Color) != 7 {
		return nil, errors.New("invalid color, format must be #hhhhhh")
	}

	// Parse the color into decimal numbers.
	colorHex, err := strconv.ParseInt((c.Color)[1:], 16, 64)
	if err != nil {
		return nil, err
	}

	// Make a new embed
	embed := discord.Embed{
		Title:       c.Title,
		Description: c.Description,
		Image:       &discord.EmbedImage{URL: c.Image},
		Color:       discord.Color(colorHex),
	}

	if c.Author != "" {
		embed.Author = &discord.EmbedAuthor{
			Name: c.Author,
		}
	}

	if c.Footer != "" {
		embed.Footer = &discord.EmbedFooter{
			Text: c.Footer,
		}
	}

	return &embed, err
}
