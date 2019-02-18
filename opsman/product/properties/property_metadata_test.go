package properties_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/opsman/product/properties"
)

var _ = Describe("PropertyMetadata", func() {
	var (
		propertyBytes    []byte
		propertyMetadata PropertyMetadata
		err              error
	)

	BeforeEach(func() {
		propertyBytes = []byte{}
		propertyMetadata = PropertyMetadata{}
	})

	Describe("MarshalJSON", func() {
		JustBeforeEach(func() {
			propertyBytes, err = propertyMetadata.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when marshaling a `multi_select_options`", func() {
			BeforeEach(func() {
				propertyMetadata.Type = PropertyTypeMultiSelectOptions
				propertyMetadata.Value.IsSet = true
				propertyMetadata.Value.Value = PropertyValueMultiSelectOptions{
					NonExistentValue: true,
				}
			})

			It("marshals the property metadata correctly", func() {
				expectedBytes := []byte(`{
"configurable": false,
"credential": false,
"optional": false,
"options": null,
"selected_option": "",
"type": "multi_select_options",
"value": "non-existant-value"
}`)
				Expect(expectedBytes).To(MatchJSON(propertyBytes))
			})
		})

		Context("when marshalling a null value", func() {
			BeforeEach(func() {
				propertyMetadata.Type = PropertyTypeInteger
				propertyMetadata.Value.IsSet = false
			})

			It("sets the property metadata value to a correct PropertyValue", func() {
				expectedBytes := []byte(`{
"configurable": false,
"credential": false,
"optional": false,
"options": null,
"selected_option": "",
"type": "integer",
"value": null
}`)
				Expect(expectedBytes).To(MatchJSON(propertyBytes))
			})
		})

		Context("when marshalling a not null value", func() {
			BeforeEach(func() {
				propertyMetadata.Type = PropertyTypeInteger
				propertyMetadata.Value.IsSet = true
				propertyMetadata.Value.Value = PropertyValueInteger(5)
			})

			It("sets the property metadata value to a correct PropertyValue", func() {
				expectedBytes := []byte(`{
"configurable": false,
"credential": false,
"optional": false,
"options": null,
"selected_option": "",
"type": "integer",
"value": 5
}`)
				Expect(expectedBytes).To(MatchJSON(propertyBytes))
			})
		})

		Context("when marshalling a collection", func() {
			BeforeEach(func() {
				propertyMetadata.Type = PropertyTypeCollection
				propertyMetadata.Value = PropertyValue{
					IsSet: true,
					Value: []map[string]PropertyMetadata{
						{
							"guid": PropertyMetadata{
								Type:  PropertyTypeUUID,
								Value: PropertyValue{IsSet: true, Value: PropertyValueUUID("beep")},
							},
							"name": PropertyMetadata{
								Type:  PropertyTypeString,
								Value: PropertyValue{IsSet: true, Value: PropertyValueString("bob")},
							},
						},
						{
							"guid": PropertyMetadata{
								Type:  PropertyTypeUUID,
								Value: PropertyValue{IsSet: true, Value: PropertyValueUUID("boop")},
							},
							"href": PropertyMetadata{
								Type:  PropertyTypeString,
								Value: PropertyValue{IsSet: true, Value: PropertyValueString("google.com")},
							},
						},
					}}
			})

			It("Marshals it into the correct json object", func() {
				expectedBytes := []byte(`{
"configurable": false,
"credential": false,
"optional": false,
"options": null,
"selected_option": "",
"type": "collection",
"value": [
	 {
		 "guid": {
			"configurable": false,
			"credential": false,
			"optional": false,
			"options": null,
			"selected_option": "",
			"type": "uuid",
			 "value": "beep"
		 },
		 "name": {
			"configurable": false,
			"credential": false,
			"optional": false,
			"options": null,
			"selected_option": "",
			"type": "string",
			"value": "bob"
		}
	 },
	 {
		 "guid": {
			"configurable": false,
			"credential": false,
			"optional": false,
			"options": null,
			"selected_option": "",
			"type": "uuid",
			"value": "boop"
		},
		 "href": {
			"configurable": false,
			"credential": false,
			"optional": false,
			"options": null,
			"selected_option": "",
		 "type": "string",
		 "value": "google.com"}
	 }
 ]
}`)

				fmt.Fprintf(GinkgoWriter, "%s", expectedBytes)
				fmt.Fprintf(GinkgoWriter, "%s", propertyBytes)
				Expect(propertyBytes).To(MatchJSON(expectedBytes))
			})
		})
	})

	Describe("UnmarshalJSON", func() {
		JustBeforeEach(func() {
			err = propertyMetadata.UnmarshalJSON(propertyBytes)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when a 'non-existant-value' is passed in", func() {
			BeforeEach(func() {
				propertyBytes = []byte(`{"type": "multi_select_options", "value": "non-existant-value"}`)
			})

			It("sets the property metadata value to a correct PropertyValue", func() {
				Expect(propertyMetadata.Type).To(Equal(PropertyTypeMultiSelectOptions))
				Expect(propertyMetadata.Value.IsSet).To(BeTrue())
				Expect(propertyMetadata.Value.Value).To(Equal(PropertyValueMultiSelectOptions{
					NonExistentValue: true,
				}))
			})
		})

		Context("when a null value is passed in", func() {
			BeforeEach(func() {
				propertyBytes = []byte(`{"type": "integer", "value": null}`)
			})

			It("sets the property metadata value to a correct PropertyValue", func() {
				Expect(propertyMetadata.Type).To(Equal(PropertyTypeInteger))
				Expect(propertyMetadata.Value.IsSet).To(BeFalse())
			})
		})

		Context("when a not null value is passed in", func() {
			BeforeEach(func() {
				propertyBytes = []byte(`{"type": "integer", "value": 5}`)
			})

			It("sets the property metadata value to a correct NullPropertyValue", func() {
				Expect(propertyMetadata.Type).To(Equal(PropertyTypeInteger))
				Expect(propertyMetadata.Value.IsSet).To(BeTrue())
				x, ok := propertyMetadata.Value.Value.(PropertyValueInteger)
				Expect(ok).To(BeTrue())
				Expect(x).To(Equal(PropertyValueInteger(5)))
			})
		})

		Context("when a collection value is passed in", func() {
			BeforeEach(func() {
				propertyBytes = []byte(`{
"type": "collection",
"value": [
	 {
		 "guid": {"type": "uuid", "value": "beep"},
		 "name": {"type": "string", "value": "bob"}
	 },
	 {
		 "guid": {"type": "uuid", "value": "boop"},
		 "href": {"type": "string", "value": "google.com"}
	 }
 ]
}`)
			})

			It("sets the property metadata value to a correct NullPropertyValue", func() {
				Expect(propertyMetadata.Type).To(Equal(PropertyTypeCollection))
				Expect(propertyMetadata.Value.IsSet).To(BeTrue())
				x, ok := propertyMetadata.Value.Value.(PropertyValueCollection)
				Expect(ok).To(BeTrue())
				Expect(x[0]).To(Equal(map[string]PropertyMetadata{
					"guid": PropertyMetadata{
						Type:  PropertyTypeUUID,
						Value: PropertyValue{IsSet: true, Value: PropertyValueUUID("beep")},
					},
					"name": PropertyMetadata{
						Type:  PropertyTypeString,
						Value: PropertyValue{IsSet: true, Value: PropertyValueString("bob")},
					},
				}))
				Expect(x[1]).To(Equal(map[string]PropertyMetadata{
					"guid": PropertyMetadata{
						Type:  PropertyTypeUUID,
						Value: PropertyValue{IsSet: true, Value: PropertyValueUUID("boop")},
					},
					"href": PropertyMetadata{
						Type:  PropertyTypeString,
						Value: PropertyValue{IsSet: true, Value: PropertyValueString("google.com")},
					},
				}))
			})
		})

		Context("when valid other property metadata bytes are passed in", func() {
			BeforeEach(func() {
				propertyBytes = []byte(`{
"configurable": true,
"credential": true,
"optional": true,
"options": [
{
	"label": "beep",
	"value": 42
},
{
	"label": "beep",
	"value": "42"
}
],
"type": "string"
}`)
			})

			It("Unmarshals them into a PropertyMetadata struct", func() {
				Expect(propertyMetadata.Configurable).To(BeTrue())
				Expect(propertyMetadata.Credential).To(BeTrue())
				Expect(propertyMetadata.Optional).To(BeTrue())
				Expect(propertyMetadata.Type).To(Equal(PropertyTypeString))
				Expect(propertyMetadata.Options).To(ConsistOf(
					Option{
						Label: "beep",
						Value: OptionValue{Type: OptionValueTypeInteger, IntegerValue: 42},
					},
					Option{
						Label: "beep",
						Value: OptionValue{Type: OptionValueTypeString, StringValue: "42"},
					},
				))
			})
		})
	})
})
