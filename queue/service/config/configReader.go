package config

import (
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	ENV_QUEUE_CONFIG_FILE    = "QUEUE_CONFIG_FILE"
	DefaultVisibilityTimeout = 30 * 1000
)

type yamlQueueObj struct {
	Order              string
	Visibility_timeout int
}

type yamlConfig struct {
	Version float32
	Queues  map[string]yamlQueueObj
}

func getConfig() *QueueConfig {
	configFileName := getConfigFileName()
	if configFileName == "" {
		return nil
	}

	yamlObj := readConfigFile(configFileName)
	return convertYamlToQueueConfig(yamlObj)
}

func getConfigFileName() string {
	configFileName := os.Getenv(ENV_QUEUE_CONFIG_FILE)
	if configFileName != "" {
		log.Printf("Config file found: %s\n", configFileName)
	}
	return configFileName
}

func readConfigFile(filename string) *yamlConfig {
	rawdata, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Err: Failed to open config file")
	}

	data := new(yamlConfig)
	err = yaml.Unmarshal(rawdata, data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func convertYamlToQueueConfig(yconfig *yamlConfig) *QueueConfig {
	version := yconfig.Version
	queues := yconfig.Queues
	config := QueueConfig{
		Version: version,
		Queues:  make([]Queue, 0),
	}
	for name, obj := range queues {
		if obj.Visibility_timeout == 0 {
			obj.Visibility_timeout = DefaultVisibilityTimeout
		}
		queue := Queue{
			Name:               name,
			Visibility_timeout: obj.Visibility_timeout,
			Order:              obj.Order,
			Messages:           make([]string, 0),
		}
		config.Queues = append(config.Queues, queue)
	}
	sort.Slice(config.Queues, func(i, j int) bool {
		return strings.Compare(config.Queues[i].Name, config.Queues[j].Name) < 0
	})
	return &config
}
