package properties

import "encoding/json"

type OptionValueType string

const (
	OptionValueTypeString  OptionValueType = "string"
	OptionValueTypeInteger OptionValueType = "integer"
)

type OptionValue struct {
	Type         OptionValueType
	IntegerValue int
	StringValue  string
}

type Option struct {
	Label string      `json:"label"`
	Value OptionValue `json:"value"`
}

func (v OptionValue) MarshalJSON() ([]byte, error) {
	if v.Type == OptionValueTypeInteger {
		return json.Marshal(v.IntegerValue)
	} else {
		return json.Marshal(v.StringValue)
	}
}

func (v *OptionValue) UnmarshalJSON(data []byte) error {
	if data[0] != '"' {
		v.Type = OptionValueTypeInteger
		return json.Unmarshal(data, &v.IntegerValue)
	} else {
		v.Type = OptionValueTypeString
		return json.Unmarshal(data, &v.StringValue)
	}
}
