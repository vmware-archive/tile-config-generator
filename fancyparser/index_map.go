package fancyparser

import "errors"

// IndexMap is an "inverted tree" of a nested map[string]inteface{}
// where the keys are leaves of the original tree, and the Values are the
// paths (represented as Indexes) to those leaves
type IndexMap map[string][]Index

// Index represents how to index a tree where nodes are sometimes indexed by
// their key, and sometimes by their array index.
type Index struct {
	Type      IndexType
	ListIndex int
	MapIndex  string
}

// IndexType specifies whether we're indexing into a dictionary or an array
type IndexType string

const (
	IndexTypeList IndexType = "list"
	IndexTypeMap  IndexType = "map"
)

// GetPlaceholderValueIndexes returns a sub map consisting of keys from map i
// containing "placeholder" values e.g. ((parameter-name))
func (i IndexMap) GetPlaceholderValueIndexes() IndexMap {
	newIndexMap := make(IndexMap)
	for key, value := range i {
		if IsPlaceholder(key) {
			newKey := key[2 : len(key)-2]
			newIndexMap[newKey] = value
		}
	}

	return newIndexMap
}

// GetPlaceholderValueIndexes returns a sub map consisting of keys from map i
// containing actual values e.g. 1024mb
func (i IndexMap) GetHardcodedValueIndexes() IndexMap {
	newIndexMap := make(IndexMap)
	for key, value := range i {
		if !IsPlaceholder(key) {
			newIndexMap[key] = value
		}
	}

	return newIndexMap
}

// IsPlaceholder checks if the provided string starts and ends with "((" "))" respectively
func IsPlaceholder(key string) bool {
	return len(key) > 4 && key[0:2] == "((" && key[len(key)-2:] == "))"
}

// GetMapWithPrependedIndex returns a map with the provided index prepended
// to every key's index in map i
func (i IndexMap) GetMapWithPrependedIndex(index Index) IndexMap {
	newIndexMap := make(IndexMap)
	for key, indexes := range i {
		newIndexes := append([]Index{index}, indexes...)
		newIndexMap[key] = newIndexes
	}

	return newIndexMap
}

// GetPropertiesIndexMap traverses the provided nested properties and compiles
// an "inverse tree" map, where keys are the original values from the properties
// tree, and values are a list of indexes explaining how to traverse the tree
// to reach the value.
func GetPropertiesIndexMap(property interface{}) IndexMap {
	indexMap := make(IndexMap)
	switch property.(type) {
	case string:
		indexMap[property.(string)] = []Index{}
		return indexMap
	case map[string]interface{}:
		value := property.(map[string]interface{})
		for p := range value {
			ixa := GetPropertiesIndexMap(value[p])
			for ixb := range ixa {
				newIndex := Index{
					Type:     IndexTypeMap,
					MapIndex: p,
				}
				indexMap[ixb] = append([]Index{newIndex}, ixa[ixb]...)
			}
		}
	case []interface{}:
		value := property.([]interface{})
		for i, p := range value {
			ixa := GetPropertiesIndexMap(p)
			for ixb := range ixa {
				newIndex := Index{
					Type:      IndexTypeList,
					ListIndex: i,
				}
				indexMap[ixb] = append([]Index{newIndex}, ixa[ixb]...)
			}
		}
	}
	return indexMap
}

func LookupPropertyWithIndexList(indexList []Index, property interface{}) (interface{}, error) {
	switch len(indexList) {
	case 0:
		switch property.(type) {
		case bool, float64, string, nil:
			return property, nil
		default:
			return nil, NoValueAtEndOfIndexError{
				RemainingValue: property,
			}
		}
	default:
		index := indexList[0]
		switch property.(type) {
		case map[string]interface{}:
			switch index.Type {
			case IndexTypeMap:
				propertyMap := property.(map[string]interface{})
				if nestedValue, ok := propertyMap[index.MapIndex]; ok {
					return LookupPropertyWithIndexList(indexList[1:], nestedValue)
				} else {
					return nil, errors.New("nested index not found in Map")
				}
			case IndexTypeList:
				return nil, InvalidIndexTypeError{
					ProvidedType: IndexTypeList,
					RequiredType: IndexTypeMap,
				}
			default:
				return nil, errors.New("unsupported index type!")
			}
		case []interface{}:
			switch index.Type {
			case IndexTypeMap:
				return nil, InvalidIndexTypeError{
					ProvidedType: IndexTypeMap,
					RequiredType: IndexTypeList,
				}
			case IndexTypeList:
				propertyList := property.([]interface{})
				if len(propertyList) > index.ListIndex {
					return LookupPropertyWithIndexList(indexList[1:], propertyList[index.ListIndex])
				} else {
					return nil, errors.New("nested index not found in List")
				}
			default:
				return nil, errors.New("unsupported index type!")
			}
		default:
			return nil, errors.New("reached end of tree!")
		}
	}
}
