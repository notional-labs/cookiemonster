package osmosis

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	ChainId string
	Node    string
)

func GetChainIdAndNodeFromFile() {
	// Open our jsonFile
	jsonFile, _ := os.Open("../osmosis_config.json")
	// if we os.Open returns an error then handle it

	// defer the closing of our jsonFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)

	ChainId = result["ChainId"]
	Node = result["Node"]
}
