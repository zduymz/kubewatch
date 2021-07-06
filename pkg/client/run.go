
package client

import (
	"log"

	"github.com/zduymz/kubewatch/config"
	"github.com/zduymz/kubewatch/pkg/controller"
	"github.com/zduymz/kubewatch/pkg/handlers"
	"github.com/zduymz/kubewatch/pkg/handlers/slack"
)

// Run runs the event loop processing with given handler
func Run(conf *config.Config) {

	var eventHandler = ParseEventHandler(conf)
	controller.Start(conf, eventHandler)
}

// ParseEventHandler returns the respective handler object specified in the config file.
func ParseEventHandler(conf *config.Config) handlers.Handler {

	var eventHandler handlers.Handler
	switch {
	case len(conf.Handler.Slack.Channel) > 0 || len(conf.Handler.Slack.Token) > 0:
		eventHandler = new(slack.Slack)
	default:
		eventHandler = new(handlers.Default)
	}
	if err := eventHandler.Init(conf); err != nil {
		log.Fatal(err)
	}
	return eventHandler
}
