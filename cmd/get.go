package cmd

import (
	"fmt"
	"os"

	sdk "github.com/lorenzodisidoro/avocado-sdk"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "Print the stored keys",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runEGet(cfgFile)
		return err
	},
}

func runEGet(configPath string) error {
	storage, config, err := getStorageAndConfig(configPath)
	if err != nil {
		return err
	}

	avocado := sdk.Avocado{}
	err = avocado.New(storage)
	if err != nil {
		return err
	}

	keys, err := avocado.GetAllKeys()
	if err != nil {
		return err
	}

	// build ASCII table content
	var data [][]string
	for i := 0; i < len(keys); i++ {
		var element []string
		element = append(element, fmt.Sprint(i))
		element = append(element, string(keys[i]))
		element = append(element, config.Bolt.Path)
		element = append(element, config.Public)
		element = append(element, config.Bolt.Bucket)
		data = append(data, element)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Key", "Public key", "Database", "Bucket"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}
