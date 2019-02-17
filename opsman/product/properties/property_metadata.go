package properties

import (
	"encoding/json"
	"errors"
)

type PropertyMetadata struct {
	Configurable   bool          `json:"configurable"`
	Credential     bool          `json:"credential"`
	Optional       bool          `json:"optional"`
	Options        []Option      `json:"options"`
	SelectedOption string        `json:"selected_option"` // used only when the type is "selector"
	Type           PropertyType  `json:"type"`
	Value          PropertyValue `json:"value"`
}

func (p PropertyMetadata) MarshalJSON() ([]byte, error) {
	var alias struct {
		Configurable   bool         `json:"configurable"`
		Credential     bool         `json:"credential"`
		Optional       bool         `json:"optional"`
		Options        []Option     `json:"options"`
		SelectedOption string       `json:"selected_option"`
		Type           PropertyType `json:"type"`
		Value          interface{}  `json:"value"`
	}

	if p.Type == "" {
		return nil, errors.New("Can't marshal a property without knowing it's type")
	}

	alias.Configurable = p.Configurable
	alias.Credential = p.Credential
	alias.Optional = p.Optional
	alias.Options = p.Options
	alias.SelectedOption = p.SelectedOption
	alias.Type = p.Type

	if p.Value.IsSet {
		alias.Value = p.Value.Value
	}

	return json.Marshal(alias)
}

func (p *PropertyMetadata) UnmarshalJSON(data []byte) error {
	var alias struct {
		Configurable   bool         `json:"configurable"`
		Credential     bool         `json:"credential"`
		Optional       bool         `json:"optional"`
		Options        []Option     `json:"options"`
		SelectedOption string       `json:"selected_option"`
		Type           PropertyType `json:"type"`
		Value          interface{}  `json:"value"`
	}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	if alias.Type == "" {
		return errors.New("Can't unmarshal property without knowing it's type")
	}

	p.Configurable = alias.Configurable
	p.Credential = alias.Credential
	p.Optional = alias.Optional
	p.Options = alias.Options
	p.SelectedOption = alias.SelectedOption
	p.Type = alias.Type

	if alias.Value == nil {
		return nil
	} else {
		p.Value.IsSet = true
	}

	switch p.Type {
	case PropertyTypeBoolean:
		p.Value.Value = PropertyValueBoolean(alias.Value.(bool))
	case PropertyTypeCACertificate:
		p.Value.Value = PropertyValueCACertificate(alias.Value.(string))
	case PropertyTypeCollection:
		collection := PropertyValueCollection{}
		collectionBytes, err := json.Marshal(alias.Value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(collectionBytes, &collection)
		if err != nil {
			return err
		}
		p.Value.Value = collection
	case PropertyTypeDropdownSelect:
		p.Value.Value = PropertyValueDropDownSelect(alias.Value.(string))
	case PropertyTypeDiskTypeDropdown:
		p.Value.Value = PropertyValueDiskTypeDropdown(alias.Value.(string))
	case PropertyTypeEmail:
		p.Value.Value = PropertyValueEmail(alias.Value.(string))
	case PropertyTypeHTTPURL:
		p.Value.Value = PropertyValueHTTPURL(alias.Value.(string))
	case PropertyTypeInteger:
		p.Value.Value = PropertyValueInteger(alias.Value.(float64))
	case PropertyTypeIPRanges:
		p.Value.Value = PropertyValueIPRanges(alias.Value.(string))
	case PropertyTypeLDAPURL:
		p.Value.Value = PropertyValueLDAPURL(alias.Value.(string))
	case PropertyTypeMultiSelectOptions:
		options := PropertyValueMultiSelectOptions{}
		optionsBytes, err := json.Marshal(alias.Value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(optionsBytes, &options)
		if err != nil {
			return err
		}
		p.Value.Value = options
	case PropertyTypeNetworkAddress:
		p.Value.Value = PropertyValueNetworkAddress(alias.Value.(string))
	case PropertyTypePort:
		p.Value.Value = PropertyValuePort(alias.Value.(float64))
	case PropertyTypeRSACertCredentials:
		creds := PropertyValueRSACertCredentials{}
		credsBytes, err := json.Marshal(alias.Value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(credsBytes, &creds)
		if err != nil {
			return err
		}
		p.Value.Value = creds
	case PropertyTypeRSAPKeyCredentials:
		creds := PropertyValueRSAPKeyCredentials{}
		credsBytes, err := json.Marshal(alias.Value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(credsBytes, &creds)
		if err != nil {
			return err
		}
		p.Value.Value = creds
	case PropertyTypeSaltedCredentials:
		creds := PropertyValueSaltedCredentials{}
		credsBytes, err := json.Marshal(alias.Value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(credsBytes, &creds)
		if err != nil {
			return err
		}
		p.Value.Value = creds
	case PropertyTypeSecret:
		secret := PropertyValueSecret{}
		secretBytes, err := json.Marshal(alias.Value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(secretBytes, &secret)
		if err != nil {
			return err
		}
		p.Value.Value = secret
	case PropertyTypeSelector:
		p.Value.Value = PropertyValueSelector(alias.Value.(string))
	case PropertyTypeServiceNetworkAZMultiSelect:
		azs := []string{}
		for _, az := range alias.Value.([]interface{}) {
			azs = append(azs, az.(string))
		}
		p.Value.Value = azs
	case PropertyTypeServiceNetworkAZSingleSelect:
		p.Value.Value = PropertyValueServiceNetworkAZSingleSelect(alias.Value.(string))
	case PropertyTypeSimpleCredentials:
		creds := PropertyValueSimpleCredentials{}
		credsBytes, err := json.Marshal(alias.Value)
		if err != nil {
			return err
		}
		err = json.Unmarshal(credsBytes, &creds)
		if err != nil {
			return err
		}
		p.Value.Value = creds
	case PropertyTypeString:
		p.Value.Value = PropertyValueString(alias.Value.(string))
	case PropertyTypeStringList:
		p.Value.Value = PropertyValueStringList(alias.Value.(string))
	case PropertyTypeText:
		p.Value.Value = PropertyValueText(alias.Value.(string))
	case PropertyTypeUUID:
		p.Value.Value = PropertyValueUUID(alias.Value.(string))
	case PropertyTypeVMTypeDropdown:
		p.Value.Value = PropertyValueVMTypeDropdown(alias.Value.(string))
	case PropertyTypeWildcardDomain:
		p.Value.Value = PropertyValueWildcardDomain(alias.Value.(string))
	}

	return nil
}
