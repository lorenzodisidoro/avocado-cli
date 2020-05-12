package cmd

import (
	"bytes"
	"fmt"

	sdk "github.com/lorenzodisidoro/avocado-sdk"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var cmdEncrypt = &cobra.Command{
	Use:   "encrypt [key]",
	Short: "Encrypt and store value, if the key ",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Expected key, value and private-key path arguments")
		}

		key := args[0]

		fmt.Print("Insert the value to encrypt: ")
		value, _ := terminal.ReadPassword(0)

		fmt.Print("\nConfirm the value to encrypt: ")
		confirmValue, _ := terminal.ReadPassword(0)

		if bytes.Compare(confirmValue, value) != 0 || len(value) == 0 {
			return fmt.Errorf("\nValue has not been confirmed")
		}

		config, err := parseConfigJSON(cfgFile)
		if err != nil {
			return err
		}

		storage, err := createStorageClient(config)
		if err != nil {
			return err
		}

		avocado := sdk.Avocado{}
		err = avocado.New(storage)
		if err != nil {
			return err
		}

		_, err = avocado.EecryptAndStoreValue([]byte(key), value, config.Public)
		if err != nil {
			return err
		}

		fmt.Println("\nEncrypted value stored successfully!")

		return nil
	},
}

func createStorageClient(config *Config) (*sdk.StorageClient, error) {
	var storage *sdk.StorageClient

	switch config.Storage {
	case "bolt":
		storage = &sdk.StorageClient{
			Bbolt: &sdk.BboltStorage{
				SotoragePath: config.Bolt.Path,
				BucketName:   config.Bolt.Bucket,
			},
		}
	case "redis":
		storage = &sdk.StorageClient{
			Redis: &sdk.RedisStorage{
				Address:  config.Redis.Address,
				Password: config.Redis.Password,
				DB:       config.Redis.Db,
			},
		}
	default:
		return storage, fmt.Errorf("Storage" + config.Storage + " not supported. Storage can be 'bolt' or 'redis'.")
	}

	return storage, nil
}
