

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	buildDate, gitCommit string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Long:  `print version`,
	Run: func(cmd *cobra.Command, args []string) {
		versionPrettyString()
	},
}

func versionPrettyString() {
	logrus.Info("gitCommit: ", gitCommit)
	logrus.Info("buildDate: ", buildDate)
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
