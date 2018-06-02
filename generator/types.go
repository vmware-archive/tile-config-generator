package generator

type Template struct {
	NetworkProperties *NetworkProperties     `yaml:"network-properties"`
	ProductProperties map[string]interface{} `yaml:"product-properties"`
	ResourceConfig    map[string]Resource    `yaml:"resource-config"`
}

type FormType struct {
	Description string     `yaml:"description"`
	Label       string     `yaml:"label"`
	Name        string     `yaml:"name"`
	Properties  []Property `yaml:"property_inputs"`
}

type PropertyMetadata struct {
	Configurable     bool               `yaml:"configurable"`
	Default          interface{}        `yaml:"default"`
	Name             string             `yaml:"name"`
	Type             string             `yaml:"type"`
	Options          []Option           `yaml:"options"`
	OptionTemplates  []OptionTemplate   `yaml:"option_templates"`
	PropertyMetadata []PropertyMetadata `yaml:"property_blueprints"`
}
type OptionTemplate struct {
	Name             string             `yaml:"name"`
	PropertyMetadata []PropertyMetadata `yaml:"property_blueprints"`
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

func (p *PropertyMetadata) IsString() bool {
	if p.Type == "dropdown_select" {
		_, ok := p.Options[0].Name.(string)
		return ok
	} else {
		return p.Type == "string" || p.Type == "text" ||
			p.Type == "ip_ranges" || p.Type == "string_list" ||
			p.Type == "network_address" || p.Type == "wildcard_domain" ||
			p.Type == "email" || p.Type == "ca_certificate" || p.Type == "http_url" || p.Type == "ldap_url" || p.Type == "service_network_az_single_select" || p.Type == "vm_type_dropdown"
	}
}
func (p *PropertyMetadata) IsInt() bool {
	if p.Type == "dropdown_select" {
		_, ok := p.Options[0].Name.(int)
		return ok
	} else {
		return p.Type == "port" || p.Type == "integer"
	}
}

func (p *PropertyMetadata) IsBool() bool {
	return p.Type == "boolean"
}

func (p *PropertyMetadata) IsSecret() bool {
	return p.Type == "secret"
}
func (p *PropertyMetadata) IsSimpleCredentials() bool {
	return p.Type == "simple_credentials"
}

func (p *PropertyMetadata) IsCollection() bool {
	return p.Type == "collection"
}
func (p *PropertyMetadata) IsSelector() bool {
	return p.Type == "selector"
}

func (p *PropertyMetadata) IsCertificate() bool {
	return p.Type == "rsa_cert_credentials"
}

func (p *PropertyMetadata) DataType() string {
	if p.IsString() {
		return "string"
	} else if p.IsInt() {
		return "int"
	} else if p.IsBool() {
		return "bool"
	} else {
		panic("Type " + p.Type + " not recongnized")
	}
}
