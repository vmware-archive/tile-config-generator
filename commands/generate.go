package commands

type Generate struct {
	PathToPivotalFile string `long:"pivotal-file-path" description:"path to pivotal file" required:"true"`
	BaseDirectory     string `long:"base-directory" description:"base directory to place generated config templates" required:"true"`
	UseFullVersion    bool   `long:"use-full-version" description:"use full version to struct names vs major/minor"`
	Version           string `long:"version" description:"version of tile"`
}

//Execute - generates config template and ops files
func (c *Generate) Execute([]string) error {
	return nil
}
