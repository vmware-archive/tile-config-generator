package generator

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func NewMetadata(fileBytes []byte) (*Metadata, error) {
	metadata := &Metadata{}
	err := yaml.Unmarshal(fileBytes, metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

type Metadata struct {
	Version          string             `yaml:"product_version"`
	FormTypes        []FormType         `yaml:"form_types"`
	PropertyMetadata []PropertyMetadata `yaml:"property_blueprints"`
	JobTypes         []JobType          `yaml:"job_types"`
}

func (m *Metadata) UsesServiceNetwork() bool {
	for _, job := range m.JobTypes {
		for _, propertyMetadata := range job.PropertyMetadata {
			if "service_network_az_single_select" == propertyMetadata.Type {
				return true
			}
		}
	}

	return false
}

func (m *Metadata) GetJob(jobName string) (*JobType, error) {
	for _, job := range m.JobTypes {
		if job.Name == jobName {
			return &job, nil
		}
	}
	return nil, fmt.Errorf("Job %s not found", jobName)
}
