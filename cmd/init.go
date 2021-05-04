package cmd

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var cmdInit = &cobra.Command{
	Use:   "init [privateKeyPath]",
	Short: "Initialize config.json file",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected RSA private key path")
		}

		privateKeyPath := args[0]

		err := runEInit(privateKeyPath, cfgFile, cfgDir)

		return err
	},
}

func runEInit(privateKeyPath, configFilePath, configFileDir string) error {
	fmt.Println(privateKeyPath)
	privateKey, err := getPrivateKeyFromFile(privateKeyPath)
	if err != nil {
		return err
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	err = initConfiguration(configFileDir, configFilePath)
	if err != nil {
		return err
	}

	publicKeyPath := configFileDir + "/" + defaultPublicKeyName
	ioutil.WriteFile(publicKeyPath, pubBytes, 0644)

	fmt.Println("Created RSA public key: " + publicKeyPath)
	fmt.Println("File successfully created, avocado is ready 🥑")

	return nil
}

// bytesToPrivateKey serialises bytes to rsa.PrivateKey
func bytesToPrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error

	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}

	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// getPrivateKeyFromFile return private key in rsa.PublicKey format
func getPrivateKeyFromFile(privateKeyPath string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := bytesToPrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
