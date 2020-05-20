package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	sdk "github.com/lorenzodisidoro/avocado-sdk"
	"github.com/spf13/cobra"
)

var cmdDecrypt = &cobra.Command{
	Use:   "decrypt [key] [privateKeyPath]",
	Short: "Encrypt and store value, if the key ",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("Expected key and RSA private key path")
		}

		key := args[0]
		privateKeyPath := args[1]

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

		decryptedValueBytes, err := avocado.FindAndDecryptValueBy([]byte(key), privateKeyPath)
		if err != nil {
			return err
		}

		clipboard.WriteAll(string(decryptedValueBytes))

		return nil
	},
}
