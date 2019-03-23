package fancyparser

type OpsFileMap map[string]map[string]interface{}

type PCFAutomationConfiguration struct {
	Features           OpsFileMap
	Optional           OpsFileMap
	ProductProperties  map[string]interface{}
	ResourceProperties map[string]interface{}
	ErrandProperties   map[string]interface{}
	VarsFiles          []string
}
