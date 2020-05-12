package cmd

import (
	"testing"
)

func TestReadConfigJSON(t *testing.T) {
	path := "./config_test.json"

	config, err := parseConfigJSON(path)
	if err != nil {
		t.Fatal(err)
	}

	if config.Storage != "bolt" {
		t.Fatal("Should be expected storage field equals to 'bolt'")
	}

	if config.Bolt.Bucket != "avocado" {
		t.Fatal("Should be expected bolt.bucket field equals to 'mybucket'")
	}
}
