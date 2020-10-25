package software

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ParseJSON parses JSON from file in path
func ParseJSON(path string) (Software, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return Software{}, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Software{}, err
	}

	var software Software
	err = json.Unmarshal(byteValue, &software)
	return software, err
}
