package config

import (
	"os"
	"testing"
)

const (
	exampleConfig           = "../example_config.yml"
	logExpectVersionMessage = "Expected version to be 1"
)

func TestGetConfigFileName(t *testing.T) {
	os.Setenv(ENV_QUEUE_CONFIG_FILE, exampleConfig)
	name := getConfigFileName()
	if name != exampleConfig {
		t.Logf("expected to retrieve %s\n", exampleConfig)
		t.Fail()
	}
	os.Unsetenv(ENV_QUEUE_CONFIG_FILE)
}

func TestReadConfigFile(t *testing.T) {
	yconfig := readConfigFile(exampleConfig)
	if yconfig.Version != 1 {
		t.Log(logExpectVersionMessage)
		t.Fail()
	}
	if len(yconfig.Queues) != 2 {
		t.Log("Expected to read 2 queues")
		t.Fail()
	}
	if yconfig.Queues["myqueue"].Order != "fifo" {
		t.Log("Expected queue 'myqueue' to have order of 'fifo'")
		t.Fail()
	}
	if yconfig.Queues["myqueue"].Visibility_timeout != 5000 {
		t.Log("Expected queue 'myqueue' to have visibility timeout = 5000")
		t.Fail()
	}
}

func TestConvertYamlToQueueConfig(t *testing.T) {
	yconfig := yamlConfig{
		Version: 1,
		Queues:  make(map[string]yamlQueueObj),
	}
	yconfig.Queues["myqueue"] = yamlQueueObj{
		Order:              "fifo",
		Visibility_timeout: 10,
	}
	actual := convertYamlToQueueConfig(&yconfig)
	if actual.Version != 1 {
		t.Log(logExpectVersionMessage)
		t.Fail()
	}
	if len(actual.Queues) != 1 {
		t.Log("Expected 1 queue")
		t.Fail()
	}
	if actual.Queues[0].Name != "myqueue" ||
		actual.Queues[0].Visibility_timeout != 10 ||
		actual.Queues[0].Order != "fifo" {
		t.Log("queue 'myqueue' does not match provided yaml config")
		t.Fail()
	}
}

func TestGetConfig(t *testing.T) {
	os.Setenv(ENV_QUEUE_CONFIG_FILE, exampleConfig)
	actual := getConfig()
	if actual == nil {
		t.Log("Failed to load config file")
		t.Fail()
		return
	}
	if actual.Version != 1 {
		t.Log(logExpectVersionMessage)
		t.Fail()
	}
	if len(actual.Queues) != 2 {
		t.Log("Expected 2 queues")
		t.Fail()
	}
	if actual.Queues[0].Name != "myqueue" ||
		actual.Queues[0].Visibility_timeout != 5000 ||
		actual.Queues[0].Order != "fifo" {
		t.Log("queue 'myqueue' does not match provided yaml config")
		t.Fail()
	}
	if actual.Queues[1].Name != "mystackqueue" ||
		actual.Queues[1].Visibility_timeout != DefaultVisibilityTimeout ||
		actual.Queues[1].Order != "filo" {
		t.Log("queue 'mystackqueue' does not match provided yaml config")
		t.Fail()
	}
}
