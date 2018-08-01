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
		if !strings.Contains(job.Name, ".") && job.InstanceDefinition.Configurable {
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
		if !strings.Contains(job.Name, ".") && job.InstanceDefinition.Configurable {
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

func CreateResourceOpsFiles(metadata *Metadata) (map[string][]Ops, error) {
	opsFiles := make(map[string][]Ops)
	for _, job := range metadata.JobTypes {
		if !strings.Contains(job.Name, ".") {
			AddResourceOpsFiles(determineJobName(job.Name), &job, opsFiles)
		}
	}
	return opsFiles, nil
}

func AddResourceOpsFiles(jobName string, job jobtype, opsFiles map[string][]Ops) {
	opsFiles[fmt.Sprintf("%s_elb_names", jobName)] = []Ops{
		Ops{
			Type:  "replace",
			Path:  fmt.Sprintf("/resource-config/%s?/elb_names?", jobName),
			Value: StringOpsValue(fmt.Sprintf("((%s_elb_names))", jobName)),
		},
	}
	opsFiles[fmt.Sprintf("%s_internet_connected", jobName)] = []Ops{
		Ops{
			Type:  "replace",
			Path:  fmt.Sprintf("/resource-config/%s?/internet_connected?", jobName),
			Value: StringOpsValue(fmt.Sprintf("((%s_internet_connected))", jobName)),
		},
	}
	opsFiles[fmt.Sprintf("%s_additional_vm_extensions", jobName)] = []Ops{
		Ops{
			Type:  "replace",
			Path:  fmt.Sprintf("/resource-config/%s?/additional_vm_extensions?", jobName),
			Value: StringOpsValue(fmt.Sprintf("((%s_additional_vm_extensions))", jobName)),
		},
	}
}
