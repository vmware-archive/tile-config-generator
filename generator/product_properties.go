package generator

import (
	"fmt"
	"strings"
)

func CreateProductProperties(metadata *Metadata) (map[string]interface{}, error) {
	productProperties := make(map[string]interface{})
	for _, property := range metadata.Properties() {
		propertyMetadata, err := metadata.GetPropertyMetadata(property.Reference)
		if err != nil {
			return nil, err
		}
		if propertyMetadata.Configurable && !propertyMetadata.Optional {
			if propertyMetadata.IsCollection() {
				if propertyMetadata.IsRequiredCollection() {
					productProperties[property.Reference] = propertyMetadata.CollectionPropertyType(strings.Replace(property.Reference, ".", "", 1))
				}
			} else {
				productProperties[property.Reference] = propertyMetadata.PropertyType(strings.Replace(property.Reference, ".", "", 1))
			}
		}

		if propertyMetadata.IsSelector() {
			defaultSelector := fmt.Sprintf("%s.%s", property.Reference, propertyMetadata.Default)
			for _, selector := range property.Selectors {
				if defaultSelector == selector.Reference {
					selectorMetadata, err := propertyMetadata.SelectorMetadata(fmt.Sprintf("%s", propertyMetadata.Default))
					if err != nil {
						return nil, err
					}
					for _, metadata := range selectorMetadata {
						selectorProperty := fmt.Sprintf("%s.%s", selector.Reference, metadata.Name)
						productProperties[selectorProperty] = metadata.PropertyType(strings.Replace(selectorProperty, ".", "", 1))
					}
				}
			}
		}
	}
	return productProperties, nil
}

func CreateProductPropertiesVars(metadata *Metadata) (map[string]interface{}, error) {
	vars := make(map[string]interface{})
	for _, property := range metadata.Properties() {
		propertyMetadata, err := metadata.GetPropertyMetadata(property.Reference)
		if err != nil {
			return nil, err
		}
		if propertyMetadata.Configurable && !propertyMetadata.Optional {
			if propertyMetadata.IsCollection() {
				if propertyMetadata.IsRequiredCollection() {
					//addPropertyToVars(property.Reference, propertyMetadata, vars)
				}
			} else {
				if !propertyMetadata.IsSelector() {
					addPropertyToVars(property.Reference, propertyMetadata, vars)
				}
			}
		}

		if propertyMetadata.IsSelector() {
			defaultSelector := fmt.Sprintf("%s.%s", property.Reference, propertyMetadata.Default)
			for _, selector := range property.Selectors {
				if defaultSelector == selector.Reference {
					selectorMetadata, err := propertyMetadata.SelectorMetadata(fmt.Sprintf("%s", propertyMetadata.Default))
					if err != nil {
						return nil, err
					}
					for _, metadata := range selectorMetadata {
						selectorProperty := fmt.Sprintf("%s.%s", selector.Reference, metadata.Name)
						addPropertyToVars(selectorProperty, &metadata, vars)
					}
				}
			}
		}
	}
	return vars, nil
}

func addPropertyToVars(propertyName string, propertyMetadata *PropertyMetadata, vars map[string]interface{}) {
	if !propertyMetadata.IsSecret() {
		newPropertyName := strings.Replace(propertyName, ".", "", 1)
		newPropertyName = strings.Replace(newPropertyName, ".", "__", -1)
		var propertyDefault interface{}
		propertyDefault = ``
		if propertyMetadata.Default != nil {
			propertyDefault = propertyMetadata.Default
		}

		vars[newPropertyName] = propertyDefault
	}
}
