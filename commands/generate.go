package commands

import (
	"github.com/pivotalservices/tile-config-generator/generator"
)

type Generate struct {
	PathToPivotalFile          string `long:"pivotal-file-path" description:"path to pivotal file" required:"true"`
	BaseDirectory              string `long:"base-directory" description:"base directory to place generated config templates" required:"true"`
	DoNotIncludeProductVersion bool   `long:"do-not-include-product-version" description:"flag to use a flat output folder"`
}

//Execute - generates config template and ops files
func (c *Generate) Execute([]string) error {
	return generator.NewExecutor(c.PathToPivotalFile, c.BaseDirectory, c.DoNotIncludeProductVersion).Generate()
}
