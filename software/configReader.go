package software

import (
	"encoding/json"
	"installscript/constants"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// ParseJSON parses JSON from file in path
func ParseJSONFile(path string) (Software, error) {
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

func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// ParseJSON Reads the configFile from the remote URL or from providedSoftware.json
func ParseJSON() (Software, error) {
	sw := &Software{}
	if getJson(constants.SOFTWAREURL, sw) != nil {
		log.Debug("Could not fetch remote Repo for provided software")
		log.Info("Loading provided software from file: ", constants.SOFTWAREFILE)
		return ParseJSONFile(constants.SOFTWAREFILE)
	}
	log.Info("Loading provided software from URL: ", constants.SOFTWAREURL)
	return *sw, nil
}
