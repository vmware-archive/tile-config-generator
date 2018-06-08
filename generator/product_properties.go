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
						if metadata.Configurable {
							selectorProperty := fmt.Sprintf("%s.%s", selector.Reference, metadata.Name)
							productProperties[selectorProperty] = metadata.PropertyType(strings.Replace(selectorProperty, ".", "", 1))
						}
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
					propertyMetadata.CollectionPropertyVars(strings.Replace(property.Reference, ".", "", 1), vars)
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
						if metadata.Configurable {
							selectorProperty := fmt.Sprintf("%s.%s", selector.Reference, metadata.Name)
							addPropertyToVars(selectorProperty, &metadata, vars)
						}
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
		newPropertyName = strings.Replace(newPropertyName, "properties.", "", 1)
		newPropertyName = strings.Replace(newPropertyName, ".", "__", -1)
		if propertyMetadata.Default != nil {
			vars[newPropertyName] = propertyMetadata.Default
		} else if propertyMetadata.IsBool() {
			vars[newPropertyName] = false
		}
	}
}

func CreateProductPropertiesOptionalOpsFiles(metadata *Metadata) (map[string][]Ops, error) {
	opsFiles := make(map[string][]Ops)
	for _, property := range metadata.Properties() {
		propertyMetadata, err := metadata.GetPropertyMetadata(property.Reference)
		if err != nil {
			return nil, err
		}
		if propertyMetadata.Configurable && propertyMetadata.Optional && !propertyMetadata.IsSelector() {

			if propertyMetadata.IsCollection() {
				for i := 1; i <= 10; i++ {
					var ops []Ops
					opsFileName := strings.Replace(property.Reference, ".", "", 1)
					opsFileName = strings.Replace(opsFileName, "properties.", "", 1)
					opsFileName = strings.Replace(opsFileName, ".", "-", -1)
					ops = append(ops,
						Ops{
							Type:  "replace",
							Path:  fmt.Sprintf("/product-properties/%s?", property.Reference),
							Value: propertyMetadata.CollectionOpsFile(i, strings.Replace(property.Reference, ".", "", 1)),
						},
					)
					opsFiles[fmt.Sprintf("add-%d-%s", i, opsFileName)] = ops
				}
			} else {
				var ops []Ops
				opsFileName := strings.Replace(property.Reference, ".", "", 1)
				opsFileName = strings.Replace(opsFileName, "properties.", "", 1)
				opsFileName = strings.Replace(opsFileName, ".", "-", -1)
				ops = append(ops,
					Ops{
						Type:  "replace",
						Path:  fmt.Sprintf("/product-properties/%s?", property.Reference),
						Value: propertyMetadata.PropertyType(strings.Replace(property.Reference, ".", "", 1)),
					},
				)
				opsFiles[fmt.Sprintf("add-%s", opsFileName)] = ops
			}

		}
	}

	return opsFiles, nil
}

func CreateProductPropertiesFeaturesOpsFiles(metadata *Metadata) (map[string][]Ops, error) {
	opsFiles := make(map[string][]Ops)
	for _, property := range metadata.Properties() {
		propertyMetadata, err := metadata.GetPropertyMetadata(property.Reference)
		if err != nil {
			return nil, err
		}

		if propertyMetadata.IsSelector() {
			defaultSelector := fmt.Sprintf("%s.%s", property.Reference, propertyMetadata.Default)
			for _, selector := range property.Selectors {
				if defaultSelector != selector.Reference {
					var ops []Ops
					opsFileName := strings.Replace(selector.Reference, ".", "", 1)
					opsFileName = strings.Replace(opsFileName, "properties.", "", 1)
					opsFileName = strings.Replace(opsFileName, ".", "-", -1)
					ops = append(ops,
						Ops{
							Type: "replace",
							Path: fmt.Sprintf("/product-properties/%s", property.Reference),
							Value: map[string]string{
								"value": strings.Replace(selector.Reference, property.Reference+".", "", 1),
							},
						},
					)

					defaultSelectorMetadata, err := propertyMetadata.SelectorMetadata(fmt.Sprintf("%s", propertyMetadata.Default))
					if err != nil {
						return nil, err
					}
					for _, metadata := range defaultSelectorMetadata {
						selectorProperty := fmt.Sprintf("%s.%s", defaultSelector, metadata.Name)
						ops = append(ops,
							Ops{
								Type: "remove",
								Path: fmt.Sprintf("/product-properties/%s?", selectorProperty),
							},
						)
					}

					selectorParts := strings.Split(selector.Reference, ".")
					selectorMetadata, err := propertyMetadata.SelectorMetadata(selectorParts[len(selectorParts)-1])
					if err != nil {
						return nil, err
					}
					for _, metadata := range selectorMetadata {
						selectorProperty := fmt.Sprintf("%s.%s", selector.Reference, metadata.Name)
						ops = append(ops,
							Ops{
								Type:  "replace",
								Path:  fmt.Sprintf("/product-properties/%s?", selectorProperty),
								Value: metadata.PropertyType(strings.Replace(selectorProperty, ".", "", 1)),
							},
						)
					}
					opsFiles[opsFileName] = ops
				}
			}
		}
	}
	return opsFiles, nil
}
