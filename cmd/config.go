package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

	file, err := json.MarshalIndent(config, "", " ")

	err = ioutil.WriteFile(configFile, file, 0644)

	return err
}
