package slack

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"

	"github.com/zduymz/kubewatch/config"
	"github.com/zduymz/kubewatch/pkg/event"
)

var slackColors = map[string]string{
	"Normal":  "good",
	"Warning": "warning",
	"Danger":  "danger",
}

var slackErrMsg = `
%s

You need to set both slack token and channel for slack notify,
using "--token/-t" and "--channel/-c", or using environment variables:

export KW_SLACK_TOKEN=slack_token
export KW_SLACK_CHANNEL=slack_channel

Command line flags will override environment variables

`

// Slack handler implements handler.Handler interface,
// Notify event to slack channel
type Slack struct {
	Token   string
	Channel string
	Title   string
}

// Init prepares slack configuration
func (s *Slack) Init(c *config.Config) error {
	token := c.Handler.Slack.Token
	channel := c.Handler.Slack.Channel
	title := c.Handler.Slack.Title

	if token == "" {
		token = os.Getenv("KW_SLACK_TOKEN")
	}

	if channel == "" {
		channel = os.Getenv("KW_SLACK_CHANNEL")
	}

	if title == "" {
		title = os.Getenv("KW_SLACK_TITLE")
		if title == "" {
			title = "kubewatch"
		}
	}

	s.Token = token
	s.Channel = channel
	s.Title = title

	return checkMissingSlackVars(s)
}

// Handle handles the notification.
func (s *Slack) Handle(e event.Event) {
	api := slack.New(s.Token)
	attachment := prepareSlackAttachment(e, s)

	channelID, timestamp, err := api.PostMessage(s.Channel,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true))
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func checkMissingSlackVars(s *Slack) error {
	if s.Token == "" || s.Channel == "" {
		return fmt.Errorf(slackErrMsg, "Missing slack token or channel")
	}

	return nil
}

func prepareSlackAttachment(e event.Event, s *Slack) slack.Attachment {

	attachment := slack.Attachment{
		Fields: []slack.AttachmentField{
			{
				Title: s.Title,
				Value: e.Message(),
			},
		},
	}

	if color, ok := slackColors[e.Status]; ok {
		attachment.Color = color
	}

	attachment.MarkdownIn = []string{"fields"}

	return attachment
}
