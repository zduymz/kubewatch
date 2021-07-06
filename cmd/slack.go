
package cmd

import (
	"github.com/zduymz/kubewatch/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// slackConfigCmd represents the slack subcommand
var slackConfigCmd = &cobra.Command{
	Use:   "slack",
	Short: "specific slack configuration",
	Long:  `specific slack configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}

		token, err := cmd.Flags().GetString("token")
		if err == nil {
			if len(token) > 0 {
				conf.Handler.Slack.Token = token
			}
		} else {
			logrus.Fatal(err)
		}
		channel, err := cmd.Flags().GetString("channel")
		if err == nil {
			if len(channel) > 0 {
				conf.Handler.Slack.Channel = channel
			}
		} else {
			logrus.Fatal(err)
		}
		title, err := cmd.Flags().GetString("title")
		if err == nil {
			if len(title) > 0 {
				conf.Handler.Slack.Title = title
			}
		}

		if err = conf.Write(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	slackConfigCmd.Flags().StringP("channel", "c", "", "Specify slack channel")
	slackConfigCmd.Flags().StringP("token", "t", "", "Specify slack token")
	slackConfigCmd.Flags().StringP("title", "", "", "Specify slack msg title")
}
