package commands

import (
	"fmt"

	"github.com/pivotalservices/tile-config-generator/fancyparser"
)

type Populate struct {
	TileConfigDirectory         string `long:"tile-config-dir" required:"true" description:"path to metadata.yml file containing awesome deets on pivotal tiles"`
	ExportedPropertiesDirectory string `long:"exported-properties-dir" required:"true" description:"path to directory containing exported tile config json files"`
	OutputFile                  string `long:"output-file" required:"true" description:"path to file where to store populate output"`
	pcfAutomationConfig         fancyparser.PCFAutomationConfiguration
	configuredProperties        fancyparser.ConfiguredProperties
	tileConfig                  fancyparser.TileConfig
}

type PropertyMetadataTest struct {
	Configurable   bool          `json:"configurable"`
	Credential     bool          `json:"credential"`
	Optional       bool          `json:"optional"`
	Options        []interface{} `json:"options"`
	SelectedOption string        `json:"selected_option"`
	Type           string        `json:"type"`
	Value          interface{}   `json:"value"`
}

// init loads up all the required files, and allocates any required maps for
// the rest of the populate command
func (c *Populate) init() error {
	configDir, err := PruneFilepath(c.TileConfigDirectory)
	if err != nil {
		return err
	}

	c.tileConfig, err = fancyparser.GetTileConfig(configDir)
	if err != nil {
		return err
	}

	propertiesDir, err := PruneFilepath(c.ExportedPropertiesDirectory)
	if err != nil {
		return err
	}

	c.configuredProperties, err = fancyparser.GetConfiguredProperties(propertiesDir)
	if err != nil {
		return err
	}

	c.pcfAutomationConfig.ProductProperties = make(map[string]interface{})
	c.pcfAutomationConfig.ResourceProperties = make(map[string]interface{})
	c.pcfAutomationConfig.ErrandProperties = make(map[string]interface{})
	c.pcfAutomationConfig.Features = make(fancyparser.OpsFileMap)
	c.pcfAutomationConfig.Optional = make(fancyparser.OpsFileMap)
	c.pcfAutomationConfig.Optional = make(fancyparser.OpsFileMap)

	return nil
}

// Execute - generates config template and ops files
func (c *Populate) Execute([]string) error {
	if err := c.init(); err != nil {
		return err
	}

	if err := c.populateProductProperties(); err != nil {
		return err
	}

	if err := c.populateOpsFiles(); err != nil {
		return err
	}

	if err := c.populateResourceProperties(); err != nil {
		return err
	}

	return writeYamlFile(c.OutputFile, c.pcfAutomationConfig)
}

func (c *Populate) populateProductProperties() error {
	indexMap := fancyparser.GetPropertiesIndexMap(c.tileConfig.ProductConfig.ProductProperties).GetPlaceholderValueIndexes()

	for placeholder, indexList := range indexMap {
		value, err := fancyparser.LookupPropertyWithRetries(indexList, c.configuredProperties.ProductProperties)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("error lookuping up: %s ", placeholder)
		}

		if value == nil {
			continue
		}

		defaultValue, ok := c.tileConfig.ProductDefaultVars[placeholder]
		if ok && defaultValue == value {
			if !itemInSlice("product-default-vars.yml", c.pcfAutomationConfig.VarsFiles) {
				c.pcfAutomationConfig.VarsFiles = append(c.pcfAutomationConfig.VarsFiles, "product-default-vars.yml")
			}
			continue
		}

		c.pcfAutomationConfig.ProductProperties[placeholder] = value
	}

	return nil
}

func (c *Populate) populateResourceProperties() error {
	indexMap := fancyparser.GetPropertiesIndexMap(c.tileConfig.ProductConfig.ResourceConfig).GetPlaceholderValueIndexes()

	for placeholder, indexList := range indexMap {
		value, err := fancyparser.LookupResourceProperty(indexList, c.configuredProperties.Resources)
		if err != nil {
			fmt.Printf("error looking up: %s", placeholder)
			fmt.Println(err)
		}

		if value == "" {
			value = "automatic"
		}

		defaultValue, ok := c.tileConfig.ResourceVars[placeholder]
		if ok && defaultValue == value {
			if !itemInSlice("resource-vars.yml", c.pcfAutomationConfig.VarsFiles) {
				c.pcfAutomationConfig.VarsFiles = append(c.pcfAutomationConfig.VarsFiles, "resource-vars.yml")
			}
			continue
		}

		c.pcfAutomationConfig.ResourceProperties[placeholder] = value
	}

	return nil
}

func (c *Populate) populateOpsFiles() error {
	for filename, opsFile := range c.tileConfig.FeaturesOpsFiles {
		err := opsFile.CheckFeatureIncludeAndGetIndexMap(c.configuredProperties.ProductProperties, c.tileConfig.ProductConfig.ProductProperties)
		if err != nil {
			return err
		}
		if opsFile.Include {
			c.pcfAutomationConfig.Features[filename] = make(map[string]interface{})
			for param, indexList := range opsFile.IndexMap {
				value, err := fancyparser.LookupPropertyWithIndexList(indexList, c.configuredProperties.ProductProperties)
				if err != nil {
					return fmt.Errorf("error lookuping up: %s from file %s", param, filename)
				}

				c.pcfAutomationConfig.Features[filename][param] = value
			}
		}
	}

	for filename, opsFile := range c.tileConfig.OptionalOpsFiles {
		err := opsFile.CheckOptionalIncludeAndGetIndexMap(c.configuredProperties.ProductProperties)
		if err != nil {
			return err
		}

		if opsFile.Include {
			c.pcfAutomationConfig.Optional[filename] = make(map[string]interface{})
			for param, indexList := range opsFile.IndexMap {
				value, err := fancyparser.LookupPropertyWithRetries(indexList, c.configuredProperties.ProductProperties)
				if err != nil {
					return fmt.Errorf("error lookuping up: %s from file %s", param, filename)
				}

				c.pcfAutomationConfig.Optional[filename][param] = value
			}
		}
	}

	return nil
}
