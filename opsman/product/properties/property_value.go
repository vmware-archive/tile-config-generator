package properties

type PropertyValue struct {
	IsSet bool
	// Type  PropertyType
	Value interface{}
}

func (v *PropertyValue) PopulateValue(value interface{}) {
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
