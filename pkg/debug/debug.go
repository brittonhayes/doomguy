package debug

import (
	"fmt"
	"github.com/diamondburned/arikawa/v2/bot/extras/middlewares"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Debug is the main struct for the
// debug subcommand
type Debug struct {
	Ctx *bot.Context
}

// NewDebug generates a new instance of Debug
func NewDebug() *Debug {
	return &Debug{}
}

// Setup initializes the debug sub command
func (d *Debug) Setup(sub *bot.Subcommand) {
	// Set a custom command (e.g. "!go ..."):
	sub.Command = "debug"
	// Set a custom description:
	sub.Description = "Debug the bot"

	// Add middleware
	sub.AddMiddleware("Die", middlewares.GuildOnly(d.Ctx))

	// Manually set the usage for each function.
	sub.ChangeCommandInfo("GOOS", "GOOS", "Prints the current operating system")
	sub.ChangeCommandInfo("GC", "GC", "Triggers the garbage collector")
	sub.ChangeCommandInfo("Goroutines", "", "Prints the current number of Goroutines")

	sub.Hide("Die")
	sub.AddMiddleware("Die", middlewares.AdminOnly(d.Ctx))
}

// Goroutines returns the number of workers running
func (d *Debug) Goroutines(*gateway.MessageCreateEvent) (string, error) {
	return fmt.Sprintf(
		"**⚡ There are currently %d workers running ⚡**",
		runtime.NumGoroutine(),
	), nil
}

// GOOS returns the operating system being used to host the bot
func (d *Debug) GOOS(*gateway.MessageCreateEvent) (string, error) {
	return strings.Title(fmt.Sprintf("🖥️ **OS:** %s", runtime.GOOS)), nil
}

// GC triggers the garbage collector
func (d *Debug) GC(*gateway.MessageCreateEvent) (string, error) {
	runtime.GC()
	return "🗑️✔️ **Garbage collected**.", nil
}

// Die kills the bot's process
func (d *Debug) Die(m *gateway.MessageCreateEvent) error {
	go func() {
		time.Sleep(5 * time.Second)
		log.Fatalf("User %s killed the bot at %v", m.Author.Username, time.Now().Format(time.Stamp))
	}()
	return fmt.Errorf("💀 %s killed the bot. \n *I will no longer respond to `+` commands*", m.Author.Username)
}
