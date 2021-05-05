package cmd

import (
	"testing"
)

func TestBytesToPrivateKey(t *testing.T) {
	privateKeyPath := "../resources/test_key.pem"

	privateKey, err := getPrivateKeyFromFile(privateKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	if privateKey.PublicKey.Size() <= 0 {
		t.Fatal("Should be expected PublicKey")
	}
}
