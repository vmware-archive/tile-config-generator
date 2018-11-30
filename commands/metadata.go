package commands

import "io/ioutil"

type MetadataCommand struct {
	PathToPivotalFile string               `long:"pivotal-file-path" description:"path to pivotal file"`
	OutputFile        string               `long:"output-file" description:"path to output metadata" required:"true"`
	Pivnet            *PivnetConfiguration `group:"pivnet"`
}

//Execute - generates config template and ops files
func (c *MetadataCommand) Execute([]string) error {
	provider, err := getProvider(c.PathToPivotalFile, c.Pivnet)
	if err != nil {
		return err
	}
	metadataBytes, err := provider.MetadataBytes()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.OutputFile, metadataBytes, 0655)
}
