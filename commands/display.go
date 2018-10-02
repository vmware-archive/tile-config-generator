package commands

import (
	"os"

	"github.com/pivotalservices/tile-config-generator/generator"
)

type Display struct {
	PathToPivotalFile string `long:"pivotal-file-path" description:"path to pivotal file" required:"true"`
}

//Execute - shows table with tile details
func (c *Display) Execute([]string) error {
	metadataBytes, err := extractMetadataBytes(c.PathToPivotalFile)
	if err != nil {
		return err
	}
	return generator.NewDisplayer(metadataBytes, os.Stdout).Display()
}
