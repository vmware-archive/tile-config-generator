package generator

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Executor struct {
	PathToPivotalFile string
	BaseDirectory     string
	ProductName       string
	Version           string
}

func NewExecutor(filePath, baseDirectory, productName, version string) *Executor {
	return &Executor{
		PathToPivotalFile: filePath,
		BaseDirectory:     baseDirectory,
		ProductName:       productName,
		Version:           version,
	}
}

func (e *Executor) Generate() error {
	metadataBytes, err := e.extractMetadataBytes()
	if err != nil {
		return err
	}
	metadata, err := NewMetadata(metadataBytes)
	if err != nil {
		return err
	}
	targetDirectory := path.Join(e.BaseDirectory, e.ProductName, e.Version)
	if err = e.createDirectory(targetDirectory); err != nil {
		return err
	}

	template, err := e.CreateTemplate(metadata)
	if err != nil {
		return err
	}

	if err = e.writeYamlFile(path.Join(targetDirectory, fmt.Sprintf("%s.yml", e.ProductName)), template); err != nil {
		return err
	}

	networkVars, err := CreateNewtworkVars(metadata)
	if err != nil {
		return err
	}

	if len(networkVars) > 0 {
		if err = e.writeYamlFile(path.Join(targetDirectory, "network-vars.yml"), networkVars); err != nil {
			return err
		}
	}

	resourceVars := CreateResourceVars(metadata)

	if len(resourceVars) > 0 {
		if err = e.writeYamlFile(path.Join(targetDirectory, "resource-vars.yml"), resourceVars); err != nil {
			return err
		}
	}

	return nil
}

func (e *Executor) CreateTemplate(metadata *Metadata) (*Template, error) {
	template := &Template{}
	template.NetworkProperties = CreateNetworkProperties(metadata)
	template.ResourceConfig = CreateResourceConfig(metadata)
	productProperties, err := CreateProductProperties(metadata)
	if err != nil {
		return nil, err
	}
	template.ProductProperties = productProperties
	return template, nil
}

func (e *Executor) createDirectory(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("cannot create directory %s: %v", path, err)
		}
	}

	return nil
}

func (e *Executor) extractMetadataBytes() ([]byte, error) {
	zipReader, err := zip.OpenReader(e.PathToPivotalFile)
	if err != nil {
		return nil, err
	}

	defer zipReader.Close()

	for _, file := range zipReader.File {
		metadataRegexp := regexp.MustCompile("metadata/.*\\.yml")
		matched := metadataRegexp.MatchString(file.Name)

		if matched {
			metadataFile, err := file.Open()
			contents, err := ioutil.ReadAll(metadataFile)
			if err != nil {
				return nil, err
			}
			return contents, nil
		}
	}
	return nil, errors.New("no metadata file was found in provided .pivotal")
}

func (e *Executor) writeYamlFile(targetFile string, dataType interface{}) error {
	data, err := yaml.Marshal(dataType)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(targetFile, data, 0755)
}
