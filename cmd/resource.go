package cmd

import (
	"github.com/zduymz/kubewatch/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resourceConfigCmd represents the resource subcommand
var resourceConfigCmd = &cobra.Command{
	Use:   "resource",
	Short: "manage resources to be watched",
	Long: `
manage resources to be watched`,
	Run: func(cmd *cobra.Command, args []string) {

		// warn for too few arguments
		if len(args) < 2 {
			logrus.Warn("Too few arguments to Command \"resource\".\nMinimum 2 arguments required: subcommand, resource flags")
		}
		// display help
		cmd.Help()
	},
}

// resourceConfigAddCmd represents the resource add subcommand
var resourceConfigAddCmd = &cobra.Command{
	Use:   "add",
	Short: "adds specific resources to be watched",
	Long: `
adds specific resources to be watched`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}

		// add resource to config
		configureResource("add", cmd, conf)
	},
}

// resourceConfigRemoveCmd represents the resource remove subcommand
var resourceConfigRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove specific resources being watched",
	Long: `
remove specific resources being watched`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}

		// remove resource from config
		configureResource("remove", cmd, conf)
	},
}

// configures resource in config based on operation add/remove
func configureResource(operation string, cmd *cobra.Command, conf *config.Config) {

	// flags struct
	flags := []struct {
		resourceStr     string
		resourceToWatch *bool
	}{
		{
			"svc",
			&conf.Resource.Services,
		},
		{
			"deploy",
			&conf.Resource.Deployment,
		},
		{
			"po",
			&conf.Resource.Pod,
		},
		{
			"rs",
			&conf.Resource.ReplicaSet,
		},
		{
			"rc",
			&conf.Resource.ReplicationController,
		},
		{
			"ns",
			&conf.Resource.Namespace,
		},
		{
			"job",
			&conf.Resource.Job,
		},
		{
			"pv",
			&conf.Resource.PersistentVolume,
		},
		{
			"ds",
			&conf.Resource.DaemonSet,
		},
		{
			"secret",
			&conf.Resource.Secret,
		},
		{
			"cm",
			&conf.Resource.ConfigMap,
		},
		{
			"ing",
			&conf.Resource.Ingress,
		},
		{
			"node",
			&conf.Resource.Node,
		},
		{
			"clusterrole",
			&conf.Resource.ClusterRole,
		},
		{
			"sa",
			&conf.Resource.ServiceAccount,
		},
	}

	for _, flag := range flags {
		b, err := cmd.Flags().GetBool(flag.resourceStr)
		if err == nil {
			if b {
				switch operation {
				case "add":
					*flag.resourceToWatch = true
					logrus.Infof("resource %s configured", flag.resourceStr)
				case "remove":
					*flag.resourceToWatch = false
					logrus.Infof("resource %s removed", flag.resourceStr)
				}
			}
		} else {
			logrus.Fatal(flag.resourceStr, err)
		}
	}

	if err := conf.Write(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	RootCmd.AddCommand(resourceConfigCmd)
	resourceConfigCmd.AddCommand(
		resourceConfigAddCmd,
		resourceConfigRemoveCmd,
	)
	// Add resource object flags as PersistentFlags to resourceConfigCmd
	resourceConfigCmd.PersistentFlags().Bool("svc", false, "watch for services")
	resourceConfigCmd.PersistentFlags().Bool("deploy", false, "watch for deployments")
	resourceConfigCmd.PersistentFlags().Bool("po", false, "watch for pods")
	resourceConfigCmd.PersistentFlags().Bool("rc", false, "watch for replication controllers")
	resourceConfigCmd.PersistentFlags().Bool("rs", false, "watch for replicasets")
	resourceConfigCmd.PersistentFlags().Bool("ns", false, "watch for namespaces")
	resourceConfigCmd.PersistentFlags().Bool("pv", false, "watch for persistent volumes")
	resourceConfigCmd.PersistentFlags().Bool("job", false, "watch for jobs")
	resourceConfigCmd.PersistentFlags().Bool("ds", false, "watch for daemonsets")
	resourceConfigCmd.PersistentFlags().Bool("secret", false, "watch for plain secrets")
	resourceConfigCmd.PersistentFlags().Bool("cm", false, "watch for plain configmaps")
	resourceConfigCmd.PersistentFlags().Bool("ing", false, "watch for ingresses")
	resourceConfigCmd.PersistentFlags().Bool("node", false, "watch for Nodes")
	resourceConfigCmd.PersistentFlags().Bool("clusterrole", false, "watch for cluster roles")
	resourceConfigCmd.PersistentFlags().Bool("sa", false, "watch for service accounts")
}
