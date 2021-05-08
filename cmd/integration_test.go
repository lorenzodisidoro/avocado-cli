package cmd

import (
	"bytes"
	"testing"

	sdk "github.com/lorenzodisidoro/avocado-sdk"
)

var (
	privateKeyPath = "../resources/test_key.pem"
	configFilePath = "../resources/config_test.json"
	configFileDir  = "../resources/"
	testKey        = "test"
	testValue      = "ciao"
)

func TestRunGetFunction(t *testing.T) {

	err := runEGet(configFilePath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRunAllFunctions(t *testing.T) {
	testValueByte := []byte(testValue)

	// init
	err := runEInit(privateKeyPath, configFilePath, configFileDir)
	if err != nil {
		t.Fatal(err)
	}

	// encrypt and decrypt new element
	err = runEEncrypt(configFilePath, testKey, testValueByte)
	if err != nil {
		t.Fatal(err)
	}

	decryptedValueBytes, err := runEDecrypt(configFilePath, testKey, privateKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(decryptedValueBytes, testValueByte) {
		t.Fatalf("The value %s was expected instead of %s", string(testValueByte), string(decryptedValueBytes))
	}

	// remove element by key
	err = runEDelete(configFilePath, testKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = runEDecrypt(configFilePath, testKey, privateKeyPath)
	if err == nil {
		t.Fatal("Was expected the error ", sdk.ErrorEncryptedValueNotFound)
	}
}
