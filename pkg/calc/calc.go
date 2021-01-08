package calc

import (
	"fmt"
	"github.com/brittonhayes/doomguy/pkg/middleware"
	"github.com/brittonhayes/doomguy/pkg/template"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	log "github.com/sirupsen/logrus"
	"strings"
)

func (c *Calc) Setup(sub *bot.Subcommand) {
	// Set a custom command (e.g. "!go ..."):
	sub.Command = "math"
	// Set a custom description:
	sub.Description = "Perform mathematical calculations"

	// Add middleware
	sub.AddMiddleware("Add", middleware.Log())
	sub.AddMiddleware("Multiply", middleware.Log())
	sub.AddMiddleware("Divide", middleware.Log())

	// Manually set the usage for each function.
	sub.ChangeCommandInfo("Add", "add", "Add up some values")
	sub.ChangeCommandInfo("Multiply", "multiply", "Multiply some values")
	sub.ChangeCommandInfo("Divide", "divide", "Divide some values")
}

// Add up all of the passed in integers
func (c *Calc) Add(_ *gateway.MessageCreateEvent, nums ...int) (string, error) {
	if len(nums) <= 0 {
		return "", fmt.Errorf(template.Usage("math add 1 2 3 4"))
	}

	result := 0
	for _, n := range nums {
		result += n
	}

	msg := fmt.Sprintf("```%s = %d```", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), " + "), "[]"), result)
	log.Info(msg)
	return msg, nil
}

// Multiply all of the passed in integers
func (c *Calc) Multiply(_ *gateway.MessageCreateEvent, a int, b int) (string, error) {
	result := a * b
	msg := fmt.Sprintf("```%v * %v = %#v```", a, b, result)
	log.Info(msg)
	return msg, nil
}

// Divide all of the passed in integers
func (c *Calc) Divide(_ *gateway.MessageCreateEvent, a int, b int) (string, error) {
	result := a / b
	msg := fmt.Sprintf("```%v / %v = %#v```", a, b, result)
	log.Info(msg)
	return msg, nil
}
