package cmd

import (
	"testing"
)

func TestRunGetFunction(t *testing.T) {
	path := "../resources/config_test.json"

	err := runEGet(path)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRunInitFunction(t *testing.T) {
	privateKeyPath := "../resources/test_key.pem"
	configFilePath := "../resources/config_test.json"
	configFileDir := "../resources/"

	err := runEInit(privateKeyPath, configFilePath, configFileDir)
	if err != nil {
		t.Fatal(err)
	}
}
