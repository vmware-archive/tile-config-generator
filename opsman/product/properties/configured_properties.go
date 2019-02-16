package properties

import "encoding/json"

type ConfiguredProperties struct {
	Properties map[string]PropertyMetadata `json:"properties"`
}

func NewConfiguredProperties(fileBytes []byte) (*ConfiguredProperties, error) {
	configuredProperties := &ConfiguredProperties{}
	err := json.Unmarshal(fileBytes, configuredProperties)
	return configuredProperties, err
}
