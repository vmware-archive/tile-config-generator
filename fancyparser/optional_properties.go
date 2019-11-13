package fancyparser

import (
	"errors"
)

// TODO: refactor. a lot of common similarities with the Features ops
func (o *OpsFile) CheckOptionalIncludeAndGetIndexMap(properties interface{}) error {
	o.IndexMap = make(IndexMap)
	for _, operation := range o.Ops {
		switch operation.Type {
		case OperationTypeReplace:
			// check what the type of the nested value is
			ov, ok := operation.Value.(map[string]interface{})
			if ok {
				switch ov["value"].(type) {
				case []interface{}:
					opsCollectionLength := len(ov["value"].([]interface{}))
					propertyName := GetPropertyNameFromPath(operation.Path)
					propertyValue := properties.(map[string]interface{})[propertyName].(map[string]interface{})["value"]

					if actualCollection, ok := propertyValue.([]interface{}); !ok || len(actualCollection) != opsCollectionLength {
						o.Include = false
						return nil
					}
					o.Include = true
					indexMap := GetPropertiesIndexMap(ov)
					topIndex := GetIndexFromPath(operation.Path)
					placeholderIndexes := indexMap.GetPlaceholderValueIndexes().GetMapWithPrependedIndex(topIndex)
					for k, v := range placeholderIndexes {
						o.IndexMap[k] = v
					}
					continue
				}

				indexMap := GetPropertiesIndexMap(operation.Value)
				topIndex := GetIndexFromPath(operation.Path)
				placeholderIndexes := indexMap.GetPlaceholderValueIndexes().GetMapWithPrependedIndex(topIndex)
				for _, indexes := range placeholderIndexes {
					configuredValue, err := LookupPropertyWithIndexList(indexes, properties)
					if err == nil && configuredValue != nil {
						o.Include = true
					} else {
						o.Include = false
						return nil
					}
				}

				for k, v := range placeholderIndexes {
					o.IndexMap[k] = v
				}
			} else {
				return errors.New("no nested map in optional ops file map")
			}
		default:
			return errors.New("unsupported operation type: " + string(operation.Type))
		}
	}

	return nil
}
