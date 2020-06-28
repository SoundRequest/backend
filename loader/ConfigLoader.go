package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/SoundRequest/OAuth2Server/structure"
)

// LoadConfig requires config path. returns DB Configuration and error
func LoadConfig(configPath string) (structure.DB, error) {
	jsonFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	var dbConf structure.DB
	err = json.Unmarshal(jsonFile, &dbConf)
	//fmt.Println(dbConf)
	return dbConf, err
}
