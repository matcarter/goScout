package configuration

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

// Config outlines common configuration values used for internal testing
type Config struct {
	APIKey string `json:"api_key"` // APIKey is the API Key used to access the Riot Games API
	Region string `json:"region"`  // Region is the region used to access the Riot Games API
}

// GetConfig reads in a json file given a file name and initializes a new Config struct
func GetConfig(filename string) (config *Config, err error) {
	logger := logrus.StandardLogger()

	// Read the file from disk
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"method":   "GetConfig",
			"filename": filename,
			"err":      err,
		}).Warn("Failed to open config file")

		return nil, err
	}

	// Parse the file into the Config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"method":   "GetConfig",
			"filename": filename,
			"err":      err,
		}).Warn("Failed to unmarshal config file")

		return nil, err
	}

	return config, err
}
