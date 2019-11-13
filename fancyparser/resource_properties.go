package fancyparser

import (
	"errors"
	"strings"
)

func LookupResourceProperty(indexList []Index, resources []interface{}) (interface{}, error) {
	if len(indexList) < 2 {
		return nil, errors.New("more than 1 index is required to lookup a resource property")
	}

	for _, resource := range resources {
		// check if map
		r, ok := resource.(map[string]interface{})
		if !ok {
			continue
		}

		// check if contains identifier and identifier matches
		identifier, ok := r["identifier"]
		if !ok || identifier != indexList[0].MapIndex {
			continue
		}

		// check deeper matches
		switch len(indexList) {
		case 2:
			for key, value := range r {
				if key == indexList[1].MapIndex {
					return value, nil
				}
			}
		case 3:
			sub_words := strings.Split(indexList[2].MapIndex, "_")
			for key, value := range r {
				if strings.Contains(key, indexList[1].MapIndex) {
					for _, word := range sub_words {
						if strings.Contains(key, word) {
							return value, nil
						}
					}
				}
			}
		}
	}
	return nil, errors.New("resource not found")
}
