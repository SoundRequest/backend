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
	jsonFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	var config structure.Config
	err = json.Unmarshal(jsonFile, &config)
	configLocal = &config
	return config, err
}

// Config from json
func Config() *structure.Config {
	return configLocal
}
