package fancyparser

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type ConfiguredProperties struct {
	ProductProperties interface{}   `json:"properties"`
	Resources         []interface{} `json:"resources"`
}

// TODO: test!!!
func GetConfiguredProperties(propertiesDir string) (ConfiguredProperties, error) {
	configuredProperties := ConfiguredProperties{}
	propertiesBytes, err := ioutil.ReadFile(filepath.Join(propertiesDir, "properties.json"))
	if err != nil {
		return ConfiguredProperties{}, err
	}

	err = json.Unmarshal(propertiesBytes, &configuredProperties)
	if err != nil {
		return ConfiguredProperties{}, err
	}

	resourcesBytes, err := ioutil.ReadFile(filepath.Join(propertiesDir, "resources.json"))
	if err != nil {
		return ConfiguredProperties{}, err
	}

	err = json.Unmarshal(resourcesBytes, &configuredProperties)
	if err != nil {
		return ConfiguredProperties{}, err
	}

	return configuredProperties, nil
}
