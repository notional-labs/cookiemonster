package mapping

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	MapFromDenomToPoolId map[string]string
)

func GetMapFromDenomToPoolId() error {
	// Open our jsonFile
	jsonFile, err := os.Open("../pool_config.json")
	if err != nil {
		return err
	}
	// if we os.Open returns an error then handle it

	// defer the closing of our jsonFile so that we can parse it later on
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)

	ChainId = result["ChainId"]
	Node = result["Node"]
	return nil
}
