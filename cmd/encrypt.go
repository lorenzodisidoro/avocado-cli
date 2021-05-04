package cmd

import (
	"bytes"
	"fmt"

	sdk "github.com/lorenzodisidoro/avocado-sdk"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var cmdEncrypt = &cobra.Command{
	Use:   "encrypt [key]",
	Short: "Encrypt and store value, if the key ",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected key, value and private-key path arguments")
		}

		key := args[0]

		fmt.Print("Insert the value to encrypt: ")
		value, _ := term.ReadPassword(0)

		fmt.Print("\nConfirm the value to encrypt: ")
		confirmValue, _ := term.ReadPassword(0)

		if !bytes.Equal(confirmValue, value) || len(value) == 0 {
			return fmt.Errorf("\nValue has not been confirmed")
		}

		storage, config, err := getStorageAndConfig(cfgFile)
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
