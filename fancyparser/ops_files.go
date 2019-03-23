package fancyparser

import (
	"github.com/ghodss/yaml"
)

type OpsFile struct {
	Include  bool
	Ops      []Operation
	IndexMap IndexMap
}

type Operation struct {
	Type  OperationType `json:"type"`
	Path  string        `json:"path"`
	Value interface{}   `json:"value"`
}

type OperationType string

const (
	OperationTypeReplace OperationType = "replace"
	OperationTypeRemove  OperationType = "remove"
)

// GetOpsFileMapFromDirBytes unmarshalls every provided filebytes into an OpsFile
func GetOpsFileMapFromDirBytes(dirBytes map[string][]byte) (map[string]OpsFile, error) {
	opsFileMap := make(map[string]OpsFile)
	for filename, fileBytes := range dirBytes {
		opsFile := OpsFile{}
		operations := []Operation{}
		err := yaml.Unmarshal(fileBytes, &operations)
		if err != nil {
			return nil, err
		}
		opsFile.Ops = operations
		opsFileMap[filename] = opsFile
	}

	return opsFileMap, nil
}
