package cmd

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	// used for flags
	cfgFile string
	cfgDir  string

	rootCmd = &cobra.Command{
		Use:     "avocado",
		Short:   "Avocado can be used to manage your passwords with RSA key",
		Version: "0.21.2",
	}

	defaultPublicKeyName = "public_key.pem"
	defaultBoltDB        = "avocado.db"
	defaultBoltBucker    = "avocado"
	defaultRedisAddress  = "localhost:6379"
	defaultRedisPassword = ""
	defaultRedisDB       = 0
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.avocado/config.json)")

	rootCmd.AddCommand(cmdInit)
	rootCmd.AddCommand(cmdEncrypt)
	rootCmd.AddCommand(cmdDecrypt)
	rootCmd.AddCommand(cmdGet)
	rootCmd.AddCommand(cmdDelete)
}

func initConfig() {
	if cfgFile == "" {
		home, _ := homedir.Expand("~")
		cfgDir = filepath.Join(home, ".avocado")
		cfgFile = filepath.Join(home, ".avocado/config.json")
	}
}
