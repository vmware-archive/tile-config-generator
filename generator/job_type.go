package generator

import (
	"fmt"
	"strings"
)

type JobType struct {
	Name                string               `yaml:"name"`
	PropertyMetadata    []PropertyMetadata   `yaml:"property_blueprints"`
	ResourceDefinitions []ResourceDefinition `yaml:"resource_definitions"`
}

type ResourceDefinition struct {
	Configurable bool        `yaml:"configurable"`
	Default      interface{} `yaml:"default"`
	Name         string      `yaml:"name"`
	Type         string      `yaml:"type"`
}

func (j *JobType) HasPersistentDisk() bool {
	for _, resourceDef := range j.ResourceDefinitions {
		if resourceDef.Name == "persistent_disk" && resourceDef.Configurable {
			return true
		}
	}
	return false
}

func (j *JobType) GetPropertyMetadata(propertyName string) (*PropertyMetadata, error) {
	propertyParts := strings.Split(propertyName, ".")
	simplePropertyName := propertyParts[len(propertyParts)-1]

	for _, property := range j.PropertyMetadata {
		if property.Name == simplePropertyName {
			return &property, nil
		}
	}
	return nil, fmt.Errorf("Property %s not found on job %s", propertyName, j.Name)
}
