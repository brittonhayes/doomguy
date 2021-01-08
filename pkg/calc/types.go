package calc

import (
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
)

var _ Calculator = (*Calc)(nil)

// Calculator contains the methods available
// to the Calc type
type Calculator interface {
	Add(m *gateway.MessageCreateEvent, nums ...int) (string, error)
	Multiply(m *gateway.MessageCreateEvent, a int, b int) (string, error)
	Divide(m *gateway.MessageCreateEvent, a int, b int) (string, error)
}

// Calc is the struct that all mathematical operations
// are bound to
type Calc struct {
	// Context must not be embedded.
	Ctx *bot.Context
}

// NewCalc generates a new instance of Calc
func NewCalc() *Calc {
	return &Calc{}
}
