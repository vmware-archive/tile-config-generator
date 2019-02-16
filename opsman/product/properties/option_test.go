package properties_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/opsman/product/properties"
)

var _ = Describe("OptionValue", func() {
	var (
		optionValueBytes []byte
		optionValue      OptionValue
		err              error
	)

	Describe("MarshalJSON", func() {
		JustBeforeEach(func() {
			optionValueBytes, err = optionValue.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the option value is an integer", func() {
			BeforeEach(func() {
				optionValue.Type = OptionValueTypeInteger
				optionValue.IntegerValue = 42
			})

			It("unmarshals the value into the correct MultiOptionType", func() {
				Expect(optionValueBytes).To(Equal([]byte(`42`)))
			})
		})

		Context("when the option value is a string", func() {
			BeforeEach(func() {
				optionValue.Type = OptionValueTypeString
				optionValue.StringValue = "42"
			})

			It("unmarshals the value into the correct MultiOptionType", func() {
				Expect(optionValueBytes).To(Equal([]byte(`"42"`)))
			})
		})
	})

	Describe("UnmarshalJSON", func() {
		JustBeforeEach(func() {
			err = optionValue.UnmarshalJSON(optionValueBytes)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the option value is an integer", func() {
			BeforeEach(func() {
				optionValueBytes = []byte(`42`)
			})

			It("unmarshals the value into the correct MultiOptionType", func() {
				Expect(optionValue.Type).To(Equal(OptionValueTypeInteger))
				Expect(optionValue.IntegerValue).To(Equal(42))
			})
		})

		Context("when the option value is a string", func() {
			BeforeEach(func() {
				optionValueBytes = []byte(`"42"`)
			})

			It("unmarshals the value into the correct MultiOptionType", func() {
				Expect(optionValue.Type).To(Equal(OptionValueTypeString))
				Expect(optionValue.StringValue).To(Equal("42"))
			})
		})
	})
})
