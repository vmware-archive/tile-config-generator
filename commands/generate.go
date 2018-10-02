package commands

import (
	"github.com/pivotalservices/tile-config-generator/generator"
)

type Generate struct {
	PathToPivotalFile          string               `long:"pivotal-file-path" description:"path to pivotal file"`
	BaseDirectory              string               `long:"base-directory" description:"base directory to place generated config templates" required:"true"`
	DoNotIncludeProductVersion bool                 `long:"do-not-include-product-version" description:"flag to use a flat output folder"`
	IncludeErrands             bool                 `long:"include-errands" description:"feature flag to include errands"`
	Pivnet                     *PivnetConfiguration `group:"pivnet"`
}

//Execute - generates config template and ops files
func (c *Generate) Execute([]string) error {

	provider := getProvider(c.PathToPivotalFile, c.Pivnet)
	metadataBytes, err := provider.MetadataBytes()
	if err != nil {
		return err
	}
	return generator.NewExecutor(metadataBytes, c.BaseDirectory, c.DoNotIncludeProductVersion, c.IncludeErrands).Generate()
}
