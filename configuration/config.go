package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	APIKey string `json:"api_key"`
	Region string `json:"region"`
}

func GetConfig(filename string) (*Config, error) {
	file, _ := ioutil.ReadFile(filename)
	config := &Config{}

	err := json.Unmarshal(file, &config)

	return config, err
}
