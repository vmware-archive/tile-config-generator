package fancyparser

import "errors"

func LookupProductProperty(indexList []Index, property interface{}) (interface{}, error) {
	if len(indexList) == 0 {
		switch property.(type) {
		case bool, float64, string:
			return property, nil
		case map[string]interface{}:
			if nestedValue, ok := property.(map[string]interface{})["value"]; ok {
				return LookupProductProperty(nil, nestedValue)
			}
			return nil, errors.New("nested value not found")
		}
	} else {
		index := indexList[0]
		switch index.Type {
		case IndexTypeMap:
			// TODO: validate property is actually indexable via map
			if nestedValue, ok := property.(map[string]interface{})[index.MapIndex]; ok {
				return LookupProductProperty(indexList[1:], nestedValue)
			} else if nestedValue, ok := property.(map[string]interface{})["value"]; ok {
				return LookupProductProperty(indexList, nestedValue)
			} else {
				return nil, errors.New("nested index not found in Map")
			}
		case IndexTypeList:
			propertyList := property.([]interface{})
			if len(propertyList) > index.ListIndex {
				return LookupProductProperty(indexList[1:], propertyList[index.ListIndex])
			} else {
				return nil, errors.New("nested index not found in List")
			}
		}
	}

	return nil, nil
}
