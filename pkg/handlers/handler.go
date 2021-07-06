

package handlers

import (
	"github.com/zduymz/kubewatch/config"
	"github.com/zduymz/kubewatch/pkg/event"
	"github.com/zduymz/kubewatch/pkg/handlers/slack"
)

// Handler is implemented by any handler.
// The Handle method is used to process event
type Handler interface {
	Init(c *config.Config) error
	Handle(e event.Event)
}

// Map maps each event handler function to a name for easily lookup
var Map = map[string]interface{}{
	"default":    &Default{},
	"slack":      &slack.Slack{},
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(c *config.Config) error {
	return nil
}

// Handle handles an event.
func (d *Default) Handle(e event.Event) {
}
