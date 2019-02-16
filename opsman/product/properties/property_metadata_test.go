package properties_test

import (
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

	Describe("MarshalJSON", func() {
	})

	Describe("UnmarshalJSON", func() {
		JustBeforeEach(func() {
			err = propertyMetadata.UnmarshalJSON(propertyBytes)
		})

		Context("when a null value is passed in", func() {
			BeforeEach(func() {
				propertyBytes = []byte(`{"type": "integer", "value": null}`)
			})

			It("sets the property metadata value to a correct PropertyValue", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(propertyMetadata.Value.IsSet).To(BeFalse())
			})
		})

		Context("when a not null value is passed in", func() {
			BeforeEach(func() {
				propertyBytes = []byte(`{"type": "integer", "value": 5}`)
			})

			It("sets the property metadata value to a correct NullPropertyValue", func() {
				Expect(err).ToNot(HaveOccurred())
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
				Expect(err).ToNot(HaveOccurred())
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
				Expect(err).ToNot(HaveOccurred())

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
