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
			return fmt.Errorf("expected key and RSA private key path")
		}

		key := args[0]
		privateKeyPath := args[1]

		decryptedValueBytes, err := runEDecrypt(cfgFile, key, privateKeyPath)
		if err != nil {
			return err
		}

		err = clipboard.WriteAll(string(decryptedValueBytes))
		if err != nil {
			return err
		}

		fmt.Println("The value has been copied to clipboard.")

		return err
	},
}

func runEDecrypt(configFilePath, key, privateKeyPath string) ([]byte, error) {
	storage, _, err := getStorageAndConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	avocado := sdk.Avocado{}
	err = avocado.New(storage)
	if err != nil {
		return nil, err
	}

	decryptedValueBytes, err := avocado.FindAndDecryptValueBy([]byte(key), privateKeyPath)

	return decryptedValueBytes, err
}
