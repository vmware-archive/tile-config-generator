package generator

import (
	"fmt"
	"strings"
)

type PropertyMetadata struct {
	Configurable     bool               `yaml:"configurable"`
	Default          interface{}        `yaml:"default"`
	Optional         bool               `yaml:"optional"`
	Name             string             `yaml:"name"`
	Type             string             `yaml:"type"`
	Options          []Option           `yaml:"options"`
	OptionTemplates  []OptionTemplate   `yaml:"option_templates"`
	PropertyMetadata []PropertyMetadata `yaml:"property_blueprints"`
}
type OptionTemplate struct {
	Name             string             `yaml:"name"`
	SelectValue      string             `yaml:"select_value"`
	PropertyMetadata []PropertyMetadata `yaml:"property_blueprints"`
}

type SimpleValue struct {
	Value interface{} `yaml:"value"`
}

type SecretValue struct {
	Value interface{} `yaml:"secret"`
}

type SimpleCredentialValue struct {
	Value interface{} `yaml:"password"`
}

type CertificateValue struct {
	CertPem        string `yaml:"cert_pem"`
	CertPrivateKey string `yaml:"private_key_pem"`
}

func (p *PropertyMetadata) SelectorMetadata(selector string) ([]PropertyMetadata, error) {
	return p.selectorMetadataByFunc(
		selector,
		func(optionTemplate OptionTemplate) string {
			return optionTemplate.Name
		})
}

// SelectorMetadataBySelectValue - uses the option template SelectValue properties of each OptionTemplate to perform the property medata selection
func (p *PropertyMetadata) SelectorMetadataBySelectValue(selector string) ([]PropertyMetadata, error) {
	return p.selectorMetadataByFunc(
		selector,
		func(optionTemplate OptionTemplate) string {
			return optionTemplate.SelectValue
		})
}

func (p *PropertyMetadata) selectorMetadataByFunc(selector string, matchFunc func(optionTemplate OptionTemplate) string) ([]PropertyMetadata, error) {
	var options []string
	for _, optionTemplate := range p.OptionTemplates {
		match := matchFunc(optionTemplate)
		options = append(options, match)

		if selector == match {
			return optionTemplate.PropertyMetadata, nil
		}
	}
	return nil, fmt.Errorf("Option template not found for selector %s options include %v", selector, options)
}

func (p *PropertyMetadata) CollectionPropertyType(propertyName string) interface{} {
	propertyName = strings.Replace(propertyName, "properties.", "", 1)
	propertyName = fmt.Sprintf("%s_0", strings.Replace(propertyName, ".", "/", -1))
	var collectionProperties []map[string]interface{}
	properties := make(map[string]interface{})
	for _, subProperty := range p.PropertyMetadata {
		if subProperty.Configurable {
			if subProperty.IsSecret() {
				properties[subProperty.Name] = &SecretValue{
					Value: fmt.Sprintf("((%s/%s))", propertyName, subProperty.Name),
				}
			} else if subProperty.IsCertificate() {
				properties[subProperty.Name] = &CertificateValue{
					CertPem:        fmt.Sprintf("((%s/%s))", propertyName, "certificate"),
					CertPrivateKey: fmt.Sprintf("((%s/%s))", propertyName, "privatekey"),
				}
			} else {
				properties[subProperty.Name] = fmt.Sprintf("((%s/%s))", propertyName, subProperty.Name)
			}
		}
	}
	collectionProperties = append(collectionProperties, properties)
	return &SimpleValue{
		Value: collectionProperties,
	}
}

func (p *PropertyMetadata) CollectionPropertyVars(propertyName string, vars map[string]interface{}) {
	propertyName = strings.Replace(propertyName, "properties.", "", 1)
	propertyName = fmt.Sprintf("%s_0", strings.Replace(propertyName, ".", "/", -1))
	for _, subProperty := range p.PropertyMetadata {
		if subProperty.Configurable {
			if !subProperty.IsSecret() && !subProperty.IsSimpleCredentials() && !subProperty.IsCertificate() {
				subPropertyName := fmt.Sprintf("%s/%s", propertyName, subProperty.Name)
				if subProperty.Default != nil {
					vars[subPropertyName] = subProperty.Default
				}
			}
		}
	}
}

func (p *PropertyMetadata) CollectionOpsFile(numOfElements int, propertyName string) interface{} {
	var collectionProperties []map[string]interface{}
	for i := 1; i <= numOfElements; i++ {
		newPropertyName := strings.Replace(propertyName, "properties.", "", 1)
		newPropertyName = fmt.Sprintf("%s_%d", strings.Replace(newPropertyName, ".", "/", -1), i-1)
		properties := make(map[string]interface{})
		for _, subProperty := range p.PropertyMetadata {
			if subProperty.IsSecret() {
				properties[subProperty.Name] = &SecretValue{
					Value: fmt.Sprintf("((%s/%s))", newPropertyName, subProperty.Name),
				}
			} else if subProperty.IsCertificate() {
				properties[subProperty.Name] = &CertificateValue{
					CertPem:        fmt.Sprintf("((%s/%s))", newPropertyName, "certificate"),
					CertPrivateKey: fmt.Sprintf("((%s/%s))", newPropertyName, "privatekey"),
				}
			} else {
				properties[subProperty.Name] = fmt.Sprintf("((%s/%s))", newPropertyName, subProperty.Name)
			}
		}
		collectionProperties = append(collectionProperties, properties)
	}
	return &SimpleValue{
		Value: collectionProperties,
	}
}

func (p *PropertyMetadata) PropertyType(propertyName string) interface{} {
	propertyName = strings.Replace(propertyName, "properties.", "", 1)
	propertyName = strings.Replace(propertyName, ".", "/", -1)
	if p.IsSelector() {
		return &SimpleValue{
			Value: fmt.Sprintf("%s", p.Default),
		}
	}
	if p.IsCertificate() {
		return &SimpleValue{
			Value: &CertificateValue{
				CertPem:        fmt.Sprintf("((%s/%s))", propertyName, "certificate"),
				CertPrivateKey: fmt.Sprintf("((%s/%s))", propertyName, "privatekey"),
			},
		}
	}
	if p.IsSecret() {
		return &SimpleValue{
			Value: &SecretValue{
				Value: fmt.Sprintf("((%s))", propertyName),
			},
		}
	}

	if p.IsSimpleCredentials() {
		return &SimpleValue{
			Value: &SimpleCredentialValue{
				Value: fmt.Sprintf("((%s))", propertyName),
			},
		}
	}
	return &SimpleValue{
		Value: fmt.Sprintf("((%s))", propertyName),
	}
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

func (p *PropertyMetadata) IsRequiredCollection() bool {
	for _, subProperty := range p.PropertyMetadata {
		if subProperty.Configurable {
			return true
		}
	}
	return false
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
