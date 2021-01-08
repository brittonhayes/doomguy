package middleware

import (
	"github.com/diamondburned/arikawa/v2/gateway"
	log "github.com/sirupsen/logrus"
)

func Log() func(interface{}) error {
	return func(ev interface{}) error {
		log.Info(ev.(*gateway.MessageCreateEvent).Content)
		return nil
	}
}
