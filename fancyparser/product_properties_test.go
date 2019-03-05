package fancyparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("LookupProductProperty", func() {
	var (
		productProperties interface{}
		indexList         []Index
		value             interface{}
		err               error
	)

	JustBeforeEach(func() {
		value, err = LookupProductProperty(indexList, productProperties)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("when the product properties contain a nested map", func() {
		BeforeEach(func() {
			indexList = []Index{
				Index{Type: IndexTypeMap, MapIndex: ".uaa.service_provider_key_credentials"},
				Index{Type: IndexTypeMap, MapIndex: "value"},
				Index{Type: IndexTypeMap, MapIndex: "cert_pem"},
			}
			productProperties = map[string]interface{}{
				".uaa.service_provider_key_credentials": map[string]interface{}{
					"value": map[string]interface{}{
						"cert_pem":        "-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n",
						"private_key_pem": "***",
					},
				},
			}
		})

		It("creates a valid IndexMap", func() {
			Expect(value).To(Equal("-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n"))
		})
	})

	Context("when the index map skips an extra 'value' nesting", func() {
		BeforeEach(func() {
			indexList = []Index{
				Index{Type: IndexTypeMap, MapIndex: ".uaa.service_provider_key_credentials"},
				Index{Type: IndexTypeMap, MapIndex: "cert_pem"},
			}
			productProperties = map[string]interface{}{
				".uaa.service_provider_key_credentials": map[string]interface{}{
					"value": map[string]interface{}{
						"cert_pem":        "-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n",
						"private_key_pem": "***",
					},
				},
			}
		})

		It("creates a valid IndexMap", func() {
			Expect(value).To(Equal("-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n"))
		})
	})

	Context("when the product properties contain a nested list", func() {
		BeforeEach(func() {
			indexList = []Index{
				Index{Type: IndexTypeMap, MapIndex: ".properties.networking_poe_ssl_certs"},
				Index{Type: IndexTypeMap, MapIndex: "value"},
				Index{Type: IndexTypeList, ListIndex: 0},
				Index{Type: IndexTypeMap, MapIndex: "certificate"},
				Index{Type: IndexTypeMap, MapIndex: "cert_pem"},
			}
			productProperties = map[string]interface{}{
				".properties.networking_poe_ssl_certs": map[string]interface{}{
					"value": []interface{}{
						map[string]interface{}{
							"certificate": map[string]interface{}{
								"value": map[string]interface{}{
									"cert_pem":        "-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n",
									"private_key_pem": "***",
								},
							},
							"name": "((networking_poe_ssl_certs_0/name))",
						},
					},
				},
			}
		})

		It("creates a valid IndexMap", func() {
			Expect(value).To(Equal("-----BEGIN CERTIFICATE-----\nbeep\n-----END CERTIFICATE-----\n"))
		})
	})
})
