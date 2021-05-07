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
			return ErrorMissingKeyAsArg
		}

		key := args[0]

		fmt.Print("Insert the value to encrypt: ")
		value, _ := term.ReadPassword(0)

		fmt.Print("\nConfirm the value to encrypt: ")
		confirmValue, _ := term.ReadPassword(0)

		if !bytes.Equal(confirmValue, value) || len(value) == 0 {
			return ErrorValuesConfirmation
		}

		err := runEEncrypt(cfgFile, key, value)

		return err
	},
}

func runEEncrypt(configFilePath, key string, value []byte) error {
	storage, config, err := getStorageAndConfig(configFilePath)
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
}
