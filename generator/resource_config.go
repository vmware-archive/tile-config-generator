package generator

import "fmt"

type Resource struct {
	InstanceType   InstanceType    `yaml:"instance_type"`
	Instances      interface{}     `yaml:"instances"`
	PersistentDisk *PersistentDisk `yaml:"persistent_disk,omitempty"`
}

type InstanceType struct {
	ID interface{} `yaml:"id"`
}
type PersistentDisk struct {
	Size interface{} `yaml:"size_mb"`
}

//go:generate counterfeiter -o ./fakes/jobtype.go --fake-name FakeJobType . jobtype
type jobtype interface {
	HasPersistentDisk() bool
}

func CreateResourceConfig(metadata *Metadata) map[string]Resource {
	resourceConfig := make(map[string]Resource)
	for _, job := range metadata.JobTypes {
		resourceConfig[job.Name] = CreateResource(job.Name, &job)
	}
	return resourceConfig
}

func CreateResource(jobName string, job jobtype) Resource {
	resource := Resource{
		Instances: fmt.Sprintf("((%s_instances))", jobName),
		InstanceType: InstanceType{
			ID: fmt.Sprintf("((%s_instance_type))", jobName),
		},
	}
	if job.HasPersistentDisk() {
		resource.PersistentDisk = &PersistentDisk{
			Size: fmt.Sprintf("((%s_persistent_disk_size))", jobName),
		}
	}
	return resource
}
