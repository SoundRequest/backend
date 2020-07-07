package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SoundRequest/backend/structure"
)

var configLocal *structure.Config

// LoadConfig requires config path. returns DB Configuration and error
func LoadConfig(configPath string) (structure.Config, error) {
	jsonFile, errFailedToReadConfig := ioutil.ReadFile(configPath)
	if errFailedToReadConfig != nil {
		fmt.Println(errFailedToReadConfig)
	}
	fmt.Println("Successfully Opened config.json")
	var config structure.Config
	errFailedToReadConfig = json.Unmarshal(jsonFile, &config)
	configLocal = &config
	return config, errFailedToReadConfig
}

// Config from json
func Config() *structure.Config {
	return configLocal
}
