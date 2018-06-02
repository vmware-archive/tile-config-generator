package generator

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
