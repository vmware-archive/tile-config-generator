package fancyparser

import (
	"errors"
	"strings"
)

// CheckFeatureIncludeAndGetIndexMap "includes" a marks a features ops file as
// "included" if and only if it modifies a hardcoded to a value specified in properties
// If it is included, it collects all the reverse index maps of the values in the ops
func (o *OpsFile) CheckFeatureIncludeAndGetIndexMap(properties interface{}, productConfig interface{}) error {
	o.IndexMap = make(IndexMap)
	for _, operation := range o.Ops {
		switch operation.Type {
		case OperationTypeReplace:
			indexMap := GetPropertiesIndexMap(operation.Value)
			topIndex := GetIndexFromPath(operation.Path)
			hardcodedIndexes := indexMap.GetHardcodedValueIndexes().GetMapWithPrependedIndex(topIndex)
			for value, indexes := range hardcodedIndexes {
				if configuredValue, err := LookupPropertyWithIndexList(indexes, properties); err == nil {
					if value == configuredValue {
						o.Include = true
					} else {
						o.Include = false
						return nil
					}
				}
			}
			placeholderIndexes := indexMap.GetPlaceholderValueIndexes().GetMapWithPrependedIndex(topIndex)
			for k, v := range placeholderIndexes {
				o.IndexMap[k] = v
			}
		case OperationTypeRemove:
			continue
		default:
			return errors.New("unsupported operation type")
		}
	}
	return nil
}

func GetIndexFromPath(path string) Index {
	return Index{
		Type:     IndexTypeMap,
		MapIndex: GetPropertyNameFromPath(path),
	}
}

func GetPropertyNameFromPath(path string) string {
	trimLeft := strings.TrimPrefix(path, "/product-properties/")
	trimRight := strings.TrimSuffix(trimLeft, "?")
	return trimRight
}
