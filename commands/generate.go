package commands

import (
	"github.com/calebwashburn/tile-config-template-generator/generator"
)

type Generate struct {
	PathToPivotalFile string `long:"pivotal-file-path" description:"path to pivotal file" required:"true"`
	BaseDirectory     string `long:"base-directory" description:"base directory to place generated config templates" required:"true"`
	ProductName       string `long:"product-name" description:"name of product" required:"true"`
	Version           string `long:"version" description:"version for path" required:"true"`
}

//Execute - generates config template and ops files
func (c *Generate) Execute([]string) error {
	return generator.NewExecutor(c.PathToPivotalFile, c.BaseDirectory, c.ProductName, c.Version).Generate()
}
