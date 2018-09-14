package generator

type Template struct {
	OriginalProductName string                   `yaml:"product_name,omitempty"` //Remove eventually added for backward compatibility
	ProductName         string                   `yaml:"product-name"`
	ProductVersion      string                   `yaml:"product_version"`
	NetworkProperties   *NetworkProperties       `yaml:"network-properties"`
	ProductProperties   map[string]PropertyValue `yaml:"product-properties"`
	ResourceConfig      map[string]Resource      `yaml:"resource-config,omitempty"`
	ErrandConfig        map[string]Errand        `yaml:"errand-config,omitempty"`
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
