package commands

import (
	"github.com/pivotalservices/tile-config-template-generator/generator"
)

type Generate struct {
	PathToPivotalFile string `long:"pivotal-file-path" description:"path to pivotal file" required:"true"`
	BaseDirectory     string `long:"base-directory" description:"base directory to place generated config templates" required:"true"`
}

//Execute - generates config template and ops files
func (c *Generate) Execute([]string) error {
	return generator.NewExecutor(c.PathToPivotalFile, c.BaseDirectory).Generate()
}
