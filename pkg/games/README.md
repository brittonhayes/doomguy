<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# games

```go
import "github.com/brittonhayes/doomguy/pkg/games"
```

## Index

- [type Gamer](<#type-gamer>)
- [type Games](<#type-games>)
  - [func NewGamesClient(t *template.Templates) *Games](<#func-newgamesclient>)
  - [func (g *Games) Get(m *gateway.MessageCreateEvent, query bot.ArgumentParts) (*discord.Embed, error)](<#func-games-get>)
  - [func (g *Games) Setup(sub *bot.Subcommand)](<#func-games-setup>)


## type Gamer

Gamer contains the methods available to the Games type

```go
type Gamer interface {
    Get(m *gateway.MessageCreateEvent, query bot.ArgumentParts) (*discord.Embed, error)
}
```

## type Games

Games contains the fields used to query the RAWG REST api and return information about video games

```go
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
```

### func NewGamesClient

```go
func NewGamesClient(t *template.Templates) *Games
```

NewGamesClient generates a new Games instance from the Games config

### func \(\*Games\) Get

```go
func (g *Games) Get(m *gateway.MessageCreateEvent, query bot.ArgumentParts) (*discord.Embed, error)
```

Get data about a specific game based on the user's query

### func \(\*Games\) Setup

```go
func (g *Games) Setup(sub *bot.Subcommand)
```

Setup demonstrates the CanSetup interface\. This function will never be parsed as a callback of any event\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
