package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	sdk "github.com/lorenzodisidoro/avocado-sdk"
)

// Config define fields of JSON file of root configurations
type Config struct {
	Storage string `json:"storage"`
	Bolt    Bolt   `json:"bolt"`
	Redis   Redis  `json:"redis"`
	Public  string `json:"publicPath"`
}

// Bolt define bolt fields
type Bolt struct {
	Path   string `json:"path"`
	Bucket string `json:"bucket"`
}

// Redis define redis fields
type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// parseConfigJSON unmarshal JSON configuration
func parseConfigJSON(configPath string) (*Config, error) {
	var configuration Config

	// open JSON file
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	byteJSON, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	json.Unmarshal(byteJSON, &configuration)

	return &configuration, nil
}

// createConfigFolder creates a new folder
func createConfigFolder(path string) error {
	err := os.Mkdir(path, 0755)
	return err
}

// exists check if the directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func initConfiguration(configDir, configFile string) error {
	config := Config{
		Storage: "bolt",
		Public:  configDir + "/" + defaultPublicKeyName,
		Bolt: Bolt{
			Path:   configDir + "/" + defaultBoltDB,
			Bucket: defaultBoltBucker,
		},
		Redis: Redis{
			Address:  defaultRedisAddress,
			Password: defaultRedisPassword,
			Db:       defaultRedisDB,
		},
	}

	configDirectoryExists, err := exists(cfgDir)
	if err != nil {
		return err
	}

	if !configDirectoryExists {
		createConfigFolder(cfgDir)
	}

	file, err := json.MarshalIndent(config, "", " ")

	err = ioutil.WriteFile(configFile, file, 0644)

	return err
}

func getStorageAndConfig(configFilePath string) (*sdk.StorageClient, *Config, error) {
	config, err := parseConfigJSON(configFilePath)
	if err != nil {
		return nil, nil, err
	}

	storage, err := createStorageClient(config)
	return storage, config, err
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
