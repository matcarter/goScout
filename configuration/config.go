package configuration

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	APIKey string `json:"api_key"`
	Region string `json:"region"`
}

func GetConfig(filename string) (config *Config, err error) {
	logger := logrus.StandardLogger()

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"method":   "GetConfig",
			"filename": filename,
			"err":      err,
		}).Warn("Failed to open config file")

		return nil, err
	}

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
