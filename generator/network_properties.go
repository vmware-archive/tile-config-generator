package generator

type NetworkProperties struct {
	Network                   *Name  `yaml:"network,omitempty"`
	ServiceNetwork            *Name  `yaml:"service_network,omitempty"`
	OtherAvailabilityZones    []Name `yaml:"other_availability_zones"`
	SingletonAvailabilityZone *Name  `yaml:"singleton_availability_zone"`
}

type Name struct {
	Name string `yaml:"name"`
}

//go:generate counterfeiter -o ./fakes/metadata.go --fake-name FakeMetadata . metadata
type metadata interface {
	UsesServiceNetwork() bool
}

func CreateNetworkProperties(metadata metadata) *NetworkProperties {
	props := &NetworkProperties{}
	props.Network = &Name{
		Name: "((network_name))",
	}
	if metadata.UsesServiceNetwork() {
		props.ServiceNetwork = &Name{
			Name: "((service_network_name))",
		}
	}
	props.SingletonAvailabilityZone = &Name{
		Name: "((singleton_availability_zone))",
	}
	props.OtherAvailabilityZones = append(props.OtherAvailabilityZones, Name{
		Name: "((singleton_availability_zone))",
	})
	return props
}

func CreateNetworkOpsFiles(metadata *Metadata) (map[string][]Ops, error) {
	opsFiles := make(map[string][]Ops)

	// var ops []Ops
	//
	// ops = append(ops,
	// 	Ops{
	// 		Type:  "replace",
	// 		Path:  fmt.Sprintf("/product-properties/%s?", property.Reference),
	// 		Value: propertyMetadata.CollectionOpsFile(i, strings.Replace(property.Reference, ".", "", 1)),
	// 	},
	// )
	opsFiles["2-az-configuration"] = []Ops{
		Ops{
			Type:  "replace",
			Path:  "/network-properties/other_availability_zones/0:after",
			Value: NameValue{Value: "((az2_name))"},
		},
	}
	opsFiles["3-az-configuration"] = []Ops{
		Ops{
			Type:  "replace",
			Path:  "/network-properties/other_availability_zones/0:after",
			Value: NameValue{Value: "((az2_name))"},
		},
		Ops{
			Type:  "replace",
			Path:  "/network-properties/other_availability_zones/1:after",
			Value: NameValue{Value: "((az3_name))"},
		},
	}

	// for _, property := range metadata.Properties() {
	// 	propertyMetadata, err := metadata.GetPropertyMetadata(property.Reference)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if propertyMetadata.Configurable && propertyMetadata.Optional && !propertyMetadata.IsSelector() {
	//
	// 		if propertyMetadata.IsCollection() {
	// 			for i := 1; i <= 10; i++ {
	// 				var ops []Ops
	// 				opsFileName := strings.Replace(property.Reference, ".", "", 1)
	// 				opsFileName = strings.Replace(opsFileName, "properties.", "", 1)
	// 				opsFileName = strings.Replace(opsFileName, ".", "-", -1)
	// 				ops = append(ops,
	// 					Ops{
	// 						Type:  "replace",
	// 						Path:  fmt.Sprintf("/product-properties/%s?", property.Reference),
	// 						Value: propertyMetadata.CollectionOpsFile(i, strings.Replace(property.Reference, ".", "", 1)),
	// 					},
	// 				)
	// 				opsFiles[fmt.Sprintf("add-%d-%s", i, opsFileName)] = ops
	// 			}
	// 		} else {
	// 			var ops []Ops
	// 			opsFileName := strings.Replace(property.Reference, ".", "", 1)
	// 			opsFileName = strings.Replace(opsFileName, "properties.", "", 1)
	// 			opsFileName = strings.Replace(opsFileName, ".", "-", -1)
	// 			ops = append(ops,
	// 				Ops{
	// 					Type:  "replace",
	// 					Path:  fmt.Sprintf("/product-properties/%s?", property.Reference),
	// 					Value: propertyMetadata.PropertyType(strings.Replace(property.Reference, ".", "", 1)),
	// 				},
	// 			)
	// 			opsFiles[fmt.Sprintf("add-%s", opsFileName)] = ops
	// 		}
	//
	// 	}
	// }

	return opsFiles, nil
}
