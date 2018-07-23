package generator

type Template struct {
	ProductName       string                   `yaml:"product_name"`
	ProductVersion    string                   `yaml:"product_version"`
	NetworkProperties *NetworkProperties       `yaml:"network-properties"`
	ProductProperties map[string]PropertyValue `yaml:"product-properties"`
	ResourceConfig    map[string]Resource      `yaml:"resource-config"`
}

type FormType struct {
	Description string     `yaml:"description"`
	Label       string     `yaml:"label"`
	Name        string     `yaml:"name"`
	Properties  []Property `yaml:"property_inputs"`
}

type Property struct {
	Description string             `yaml:"description"`
	Label       string             `yaml:"label"`
	Placeholder string             `yaml:"placeholder"`
	Reference   string             `yaml:"reference"`
	Selectors   []SelectorProperty `yaml:"selector_property_inputs"`
}

type SelectorProperty struct {
	Label      string     `yaml:"label"`
	Reference  string     `yaml:"reference"`
	Properties []Property `yaml:"property_inputs"`
}

type Option struct {
	Label string      `json:"label"`
	Name  interface{} `json:"name"`
}
