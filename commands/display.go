package commands

import (
	"os"

	"github.com/pivotalservices/tile-config-generator/generator"
)

type Display struct {
	PathToPivotalFile string               `long:"pivotal-file-path" description:"path to pivotal file"`
	Pivnet            *PivnetConfiguration `group:"pivnet"`
}

//Execute - shows table with tile details
func (c *Display) Execute([]string) error {
	provider, err := getProvider(c.PathToPivotalFile, c.Pivnet)
	if err != nil {
		return err
	}
	metadataBytes, err := provider.MetadataBytes()
	if err != nil {
		return err
	}
	return generator.NewDisplayer(metadataBytes, os.Stdout).Display()
}
