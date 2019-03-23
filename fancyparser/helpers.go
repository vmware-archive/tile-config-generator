package fancyparser

// LookupPropertyWithRetries calls LookupPropertyWithIndexList with the provided
// parameters. If an error occurs, it retries by first placing a "value" IndexMap
// at the end of the chain, then everywhere along the chain
func LookupPropertyWithRetries(indexList []Index, property interface{}) (interface{}, error) {
	foundValue, err := LookupPropertyWithIndexList(indexList, property)
	if err != nil {
		// fmt.Println(err)
		// fmt.Printf("%v", indexList)
		if newType, ok := err.(NoValueAtEndOfIndexError); ok {
			if remainingValue, ok := newType.RemainingValue.(map[string]interface{}); ok {
				if _, ok := remainingValue["value"]; ok {
					newIndex := Index{Type: IndexTypeMap, MapIndex: "value"}
					newIndexList := append(indexList, newIndex)
					return LookupPropertyWithIndexList(newIndexList, property)
				}
			}
		} else {
			l := len(indexList)
			newIndex := Index{Type: IndexTypeMap, MapIndex: "value"}
			newIndexList := append([]Index{}, indexList...)
			newIndexList = append(newIndexList, newIndex)
			copy(newIndexList[l:], newIndexList[l-1:])
			newIndexList[l-1] = newIndex

			newFoundValue, err := LookupPropertyWithIndexList(newIndexList, property)
			if err == nil {
				return newFoundValue, nil
			}
		}
		return nil, err
	}
	return foundValue, nil
}
