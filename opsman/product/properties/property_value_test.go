package properties_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/opsman/product/properties"
)

var _ = Describe("PropertyValueMultiSelectOptions", func() {
	var (
		propertyValueBytes []byte
		propertyValue      PropertyValueMultiSelectOptions
		err                error
	)

	BeforeEach(func() {
		propertyValueBytes = []byte{}
		propertyValue = PropertyValueMultiSelectOptions{}
	})

	Context("MarshalJSON", func() {
		JustBeforeEach(func() {
			propertyValueBytes, err = propertyValue.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when marshaling 'non-existant-value'", func() {
			BeforeEach(func() {
				propertyValue.NonExistentValue = true
			})

			It("sets marshals the string 'non-existant-value'", func() {
				Expect(propertyValueBytes).To(Equal([]byte(`"non-existant-value"`)))
			})
		})

		Context("when marshaling some options", func() {
			BeforeEach(func() {
				propertyValue.NonExistentValue = false
				propertyValue.Value = []string{"beep", "boop"}
			})

			It("marshals all the options into a list", func() {
				Expect(propertyValueBytes).To(Equal([]byte(`["beep","boop"]`)))
			})
		})
	})

	Context("UnmarshalJSON", func() {
		JustBeforeEach(func() {
			err = propertyValue.UnmarshalJSON(propertyValueBytes)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when unmarshaling 'non-existant-value'", func() {
			BeforeEach(func() {
				propertyValueBytes = []byte(`"non-existant-value"`)
			})

			It("sets NonExistentValue to true", func() {
				Expect(propertyValue.NonExistentValue).To(BeTrue())
			})
		})

		Context("when unmarshaling some options", func() {
			BeforeEach(func() {
				propertyValueBytes = []byte(`["beep", "boop"]`)
			})

			It("sets NonExistentValue to true", func() {
				Expect(propertyValue.NonExistentValue).To(BeFalse())
				Expect(propertyValue.Value).To(ConsistOf("beep", "boop"))
			})
		})
	})
})
