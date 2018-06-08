package generator

import (
	"fmt"
	"strings"
)

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
		if !strings.Contains(job.Name, ".") {
			resourceConfig[job.Name] = CreateResource(determineJobName(job.Name), &job)
		}
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

func CreateResourceVars(metadata *Metadata) map[string]interface{} {
	vars := make(map[string]interface{})
	for _, job := range metadata.JobTypes {
		if !strings.Contains(job.Name, ".") {
			AddResourceVars(determineJobName(job.Name), &job, vars)
		}
	}
	return vars
}

func AddResourceVars(jobName string, job jobtype, vars map[string]interface{}) {
	vars[fmt.Sprintf("%s_instances", jobName)] = "automatic"
	vars[fmt.Sprintf("%s_instance_type", jobName)] = "automatic"
	if job.HasPersistentDisk() {
		vars[fmt.Sprintf("%s_persistent_disk_size", jobName)] = "automatic"
	}
}

func determineJobName(jobName string) string {
	return strings.Replace(jobName, ".", "_", -1)
}
