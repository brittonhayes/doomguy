<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# guild

```go
import "github.com/brittonhayes/doomguy/pkg/guild"
```

## Index

- [type Guild](<#type-guild>)
  - [func NewGuild() *Guild](<#func-newguild>)
  - [func (guild *Guild) GuildInfo(m *gateway.MessageCreateEvent) (*discord.Embed, error)](<#func-guild-guildinfo>)
  - [func (guild *Guild) Setup(sub *bot.Subcommand)](<#func-guild-setup>)


## type Guild

Flag for administrators only\.

```go
type Guild struct {
    Ctx *bot.Context
}
```

### func NewGuild

```go
func NewGuild() *Guild
```

NewGuild generates a new instance of Guild

### func \(\*Guild\) GuildInfo

```go
func (guild *Guild) GuildInfo(m *gateway.MessageCreateEvent) (*discord.Embed, error)
```

GuildInfo demonstrates the GuildOnly middleware done in \(\*Bot\)\.Setup\(\)\.

### func \(\*Guild\) Setup

```go
func (guild *Guild) Setup(sub *bot.Subcommand)
```

Setup demonstrates the CanSetup interface\. This function will never be parsed as a callback of any event\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
