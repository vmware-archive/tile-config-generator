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
