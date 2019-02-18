package properties

import (
	"encoding/json"
	"errors"
	"reflect"
)

type PropertyValue struct {
	IsSet bool
	// Type  PropertyType
	Value interface{}
}

type PropertyValueBoolean bool

type PropertyValueCACertificate string

type PropertyValueCollection []map[string]PropertyMetadata

type PropertyValueDropDownSelect string

type PropertyValueDiskTypeDropdown string

type PropertyValueEmail string

type PropertyValueHTTPURL string

type PropertyValueInteger float64

type PropertyValueIPRanges string

type PropertyValueLDAPURL string

// PropertyValueMultiSelectOptions value contains either a list of strings
// or the literal string "non-existant-value"
type PropertyValueMultiSelectOptions struct {
	NonExistentValue bool
	Value            []string
}

func (p *PropertyValueMultiSelectOptions) UnmarshalJSON(propertyBytes []byte) error {
	if reflect.DeepEqual(propertyBytes, []byte(`"non-existant-value"`)) {
		p.NonExistentValue = true
	} else {
		var value interface{}
		err := json.Unmarshal(propertyBytes, &value)
		if err != nil {
			return err
		}
		switch x := value.(type) {
		case []interface{}:
			for _, option := range x {
				p.Value = append(p.Value, option.(string))
			}
		default:
			return errors.New("found unknown value in multi_select_options type")
		}
	}
	return nil
}

func (p PropertyValueMultiSelectOptions) MarshalJSON() ([]byte, error) {
	if p.NonExistentValue {
		return []byte(`"non-existant-value"`), nil
	}

	return json.Marshal(p.Value)
}

type PropertyValueNetworkAddress string

type PropertyValuePort float64

type PropertyValueRSACertCredentials struct {
	CertPem       string `json:"cert_pem"`
	PrivateKeyPem string `json:"private_key_pem"`
}

type PropertyValueRSAPKeyCredentials struct {
	PrivateKeyPem string `json:"private_key_pem"`
}

type PropertyValueSaltedCredentials struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type PropertyValueSecret struct {
	Secret string `json:"secret"`
}

type PropertyValueSelector string

type PropertyValueServiceNetworkAZMultiSelect []string

type PropertyValueServiceNetworkAZSingleSelect string

type PropertyValueSimpleCredentials struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type PropertyValueString string

type PropertyValueStringList string

type PropertyValueText string

type PropertyValueUUID string

type PropertyValueVMTypeDropdown string

type PropertyValueWildcardDomain string
