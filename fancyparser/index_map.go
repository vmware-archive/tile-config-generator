package fancyparser

type IndexMap map[string][]Index

type Index struct {
	Type      IndexType
	ListIndex int
	MapIndex  string
}

type IndexType int

const (
	IndexTypeList IndexType = iota
	IndexTypeMap
)

func GetProductPropertiesIndexMap(property interface{}) IndexMap {
	indexMap := make(IndexMap)
	switch property.(type) {
	case string:
		// fmt.Printf("it's a string: %s", property)
		value := property.(string)
		if value[0:2] == "((" && value[len(value)-2:] == "))" {
			variable := value[2 : len(value)-2]
			indexMap[variable] = []Index{}
		}
		return indexMap
	case map[string]interface{}:
		value := property.(map[string]interface{})
		for p := range value {
			ixa := GetProductPropertiesIndexMap(value[p])
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
			ixa := GetProductPropertiesIndexMap(p)
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
