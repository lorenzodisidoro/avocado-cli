package cmd

import (
	sdk "github.com/lorenzodisidoro/avocado-sdk"
	"github.com/spf13/cobra"
)

var cmdDelete = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete element by key",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return ErrorMissingKeyAsArg
		}

		key := args[0]

		err := runEDelete(cfgFile, key)

		return err
	},
}

func runEDelete(configFilePath, key string) error {

	storage, _, err := getStorageAndConfig(configFilePath)
	if err != nil {
		return err
	}

	avocado := sdk.Avocado{}
	err = avocado.New(storage)
	if err != nil {
		return err
	}

	byteKey := []byte(key)
	err = avocado.Delete(byteKey)
	if err != nil {
		return err
	}

	return nil
}
