package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

var (
	// ConfigFileName stores file of config
	ConfigFileName = ".kubewatch.yaml"

	// ConfigSample is a sample configuration file.
	ConfigSample = yannotated
)

// Handler contains handler configuration
type Handler struct {
	Slack      Slack      `json:"slack"`
}

// Resource contains resource configuration
type Resource struct {
	Deployment            bool `json:"deployment"`
	ReplicationController bool `json:"rc"`
	ReplicaSet            bool `json:"rs"`
	DaemonSet             bool `json:"ds"`
	Services              bool `json:"svc"`
	Pod                   bool `json:"po"`
	Job                   bool `json:"job"`
	Node                  bool `json:"node"`
	ClusterRole           bool `json:"clusterrole"`
	ServiceAccount        bool `json:"sa"`
	PersistentVolume      bool `json:"pv"`
	Namespace             bool `json:"ns"`
	Secret                bool `json:"secret"`
	ConfigMap             bool `json:"configmap"`
	Ingress               bool `json:"ing"`
}

// Config struct contains kubewatch configuration
type Config struct {
	// Handlers know how to send notifications to specific services.
	Handler Handler `json:"handler"`

	//Reason   []string `json:"reason"`

	// Resources to watch.
	Resource Resource `json:"resource"`

	// For watching specific namespace, leave it empty for watching all.
	// this config is ignored when watching namespaces
	Namespace string `json:"namespace,omitempty"`
}

// Slack contains slack configuration
type Slack struct {
	// Slack "legacy" API token.
	Token string `json:"token"`
	// Slack channel.
	Channel string `json:"channel"`
	// Title of the message.
	Title string `json:"title"`
}


// New creates new config object
func New() (*Config, error) {
	c := &Config{}
	if err := c.Load(); err != nil {
		return c, err
	}

	return c, nil
}

func createIfNotExist() error {
	// create file if not exist
	configFile := filepath.Join(configDir(), ConfigFileName)
	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(configFile)
			if err != nil {
				return err
			}
			file.Close()
		} else {
			return err
		}
	}
	return nil
}

// Load loads configuration from config file
func (c *Config) Load() error {
	err := createIfNotExist()
	if err != nil {
		return err
	}

	file, err := os.Open(getConfigFile())
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if len(b) != 0 {
		return yaml.Unmarshal(b, c)
	}

	return nil
}

// CheckMissingResourceEnvvars will read the environment for equivalent config variables to set
func (c *Config) CheckMissingResourceEnvvars() {
	if !c.Resource.DaemonSet && os.Getenv("KW_DAEMONSET") == "true" {
		c.Resource.DaemonSet = true
	}
	if !c.Resource.ReplicaSet && os.Getenv("KW_REPLICASET") == "true" {
		c.Resource.ReplicaSet = true
	}
	if !c.Resource.Namespace && os.Getenv("KW_NAMESPACE") == "true" {
		c.Resource.Namespace = true
	}
	if !c.Resource.Deployment && os.Getenv("KW_DEPLOYMENT") == "true" {
		c.Resource.Deployment = true
	}
	if !c.Resource.Pod && os.Getenv("KW_POD") == "true" {
		c.Resource.Pod = true
	}
	if !c.Resource.ReplicationController && os.Getenv("KW_REPLICATION_CONTROLLER") == "true" {
		c.Resource.ReplicationController = true
	}
	if !c.Resource.Services && os.Getenv("KW_SERVICE") == "true" {
		c.Resource.Services = true
	}
	if !c.Resource.Job && os.Getenv("KW_JOB") == "true" {
		c.Resource.Job = true
	}
	if !c.Resource.PersistentVolume && os.Getenv("KW_PERSISTENT_VOLUME") == "true" {
		c.Resource.PersistentVolume = true
	}
	if !c.Resource.Secret && os.Getenv("KW_SECRET") == "true" {
		c.Resource.Secret = true
	}
	if !c.Resource.ConfigMap && os.Getenv("KW_CONFIGMAP") == "true" {
		c.Resource.ConfigMap = true
	}
	if !c.Resource.Ingress && os.Getenv("KW_INGRESS") == "true" {
		c.Resource.Ingress = true
	}
	if !c.Resource.Node && os.Getenv("KW_NODE") == "true" {
		c.Resource.Node = true
	}
	if !c.Resource.ServiceAccount && os.Getenv("KW_SERVICE_ACCOUNT") == "true" {
		c.Resource.ServiceAccount = true
	}
	if !c.Resource.ClusterRole && os.Getenv("KW_CLUSTER_ROLE") == "true" {
		c.Resource.ClusterRole = true
	}
	if (c.Handler.Slack.Channel == "") && (os.Getenv("SLACK_CHANNEL") != "") {
		c.Handler.Slack.Channel = os.Getenv("SLACK_CHANNEL")
	}
	if (c.Handler.Slack.Token == "") && (os.Getenv("SLACK_TOKEN") != "") {
		c.Handler.Slack.Token = os.Getenv("SLACK_TOKEN")
	}
}

func (c *Config) Write() error {
	f, err := os.OpenFile(getConfigFile(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	enc.SetIndent(2) // compat with old versions of kubewatch
	return enc.Encode(c)
}

func getConfigFile() string {
	configFile := filepath.Join(configDir(), ConfigFileName)
	if _, err := os.Stat(configFile); err == nil {
		return configFile
	}

	return ""
}

func configDir() string {
	if configDir := os.Getenv("KW_CONFIG"); configDir != "" {
		return configDir
	}

	if runtime.GOOS == "windows" {
		home := os.Getenv("USERPROFILE")
		return home
	}
	return os.Getenv("HOME")
	//path := "/etc/kubewatch"
	//if _, err := os.Stat(path); os.IsNotExist(err) {
	//	os.Mkdir(path, 755)
	//}
	//return path
}
