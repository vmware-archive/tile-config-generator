package fancyparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("IndexMap", func() {
	Context("LookupPropertyWithIndexList", func() {
		var (
			productProperties interface{}
			indexList         []Index
			value             interface{}
			err               error
		)

		JustBeforeEach(func() {
			value, err = LookupPropertyWithRetries(indexList, productProperties)
		})

		Context("when there's a value map missing at the end", func() {
			BeforeEach(func() {
				indexList = []Index{
					Index{Type: IndexTypeMap, MapIndex: ".uaa.service_provider_key_credentials"},
					Index{Type: IndexTypeMap, MapIndex: "value"},
					Index{Type: IndexTypeMap, MapIndex: "cert_pem"},
				}

				productProperties = map[string]interface{}{
					".uaa.service_provider_key_credentials": map[string]interface{}{
						"value": map[string]interface{}{
							"cert_pem": map[string]interface{}{
								"value": "-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n",
							},
							"private_key_pem": "***",
						},
					},
				}
			})

			It("creates a valid IndexMap", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(value).To(Equal("-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n"))
			})
		})

		Context("when there's a value map missing in the second to last item", func() {
			BeforeEach(func() {
				indexList = []Index{
					Index{Type: IndexTypeMap, MapIndex: ".uaa.service_provider_key_credentials"},
					Index{Type: IndexTypeMap, MapIndex: "cert_pem"},
					// Index{Type: IndexTypeMap, MapIndex: "value"},
				}

				productProperties = map[string]interface{}{
					".uaa.service_provider_key_credentials": map[string]interface{}{
						"value": map[string]interface{}{
							"cert_pem": "beep", //map[string]interface{}{
							// "value": "-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n",
							// },
							"private_key_pem": "***",
						},
					},
				}
			})

			It("creates a valid IndexMap", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(value).To(Equal("beep"))
				// Expect(value).To(Equal("-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n"))
			})
		})
	})
})
