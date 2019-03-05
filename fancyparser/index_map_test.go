package fancyparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("IndexMap", func() {
	var (
		productProperties map[string]interface{}
		indexMap          IndexMap
		err               error
	)

	Context("GetProductPropertiesIndexMap", func() {
		JustBeforeEach(func() {
			indexMap = GetProductPropertiesIndexMap(productProperties)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the product properties contain a nested map", func() {
			BeforeEach(func() {
				productProperties = map[string]interface{}{
					".uaa.service_provider_key_credentials": map[string]interface{}{
						"value": map[string]interface{}{
							"cert_pem":        "((uaa/service_provider_key_credentials/certificate))",
							"private_key_pem": "((uaa/service_provider_key_credentials/privatekey))",
						},
					},
				}
			})

			It("creates a valid IndexMap", func() {
				Expect(indexMap).To(Equal(IndexMap{
					"uaa/service_provider_key_credentials/certificate": []Index{
						Index{Type: IndexTypeMap, MapIndex: ".uaa.service_provider_key_credentials"},
						Index{Type: IndexTypeMap, MapIndex: "value"},
						Index{Type: IndexTypeMap, MapIndex: "cert_pem"},
					},
					"uaa/service_provider_key_credentials/privatekey": []Index{
						Index{Type: IndexTypeMap, MapIndex: ".uaa.service_provider_key_credentials"},
						Index{Type: IndexTypeMap, MapIndex: "value"},
						Index{Type: IndexTypeMap, MapIndex: "private_key_pem"},
					},
				}))
			})
		})

		Context("when the product properties contain a nested slice", func() {
			BeforeEach(func() {
				productProperties = map[string]interface{}{
					".properties.networking_poe_ssl_certs": map[string]interface{}{
						"value": []interface{}{
							map[string]interface{}{
								"certificate": map[string]interface{}{
									"cert_pem":        "((networking_poe_ssl_certs_0/certificate))",
									"private_key_pem": "((networking_poe_ssl_certs_0/privatekey))",
								},
								"name": "((networking_poe_ssl_certs_0/name))",
							},
						},
					},
				}
			})

			It("creates a valid IndexMap", func() {
				Expect(indexMap).To(Equal(IndexMap{
					"networking_poe_ssl_certs_0/certificate": []Index{
						Index{Type: IndexTypeMap, MapIndex: ".properties.networking_poe_ssl_certs"},
						Index{Type: IndexTypeMap, MapIndex: "value"},
						Index{Type: IndexTypeList, ListIndex: 0},
						Index{Type: IndexTypeMap, MapIndex: "certificate"},
						Index{Type: IndexTypeMap, MapIndex: "cert_pem"},
					},
					"networking_poe_ssl_certs_0/privatekey": []Index{
						Index{Type: IndexTypeMap, MapIndex: ".properties.networking_poe_ssl_certs"},
						Index{Type: IndexTypeMap, MapIndex: "value"},
						Index{Type: IndexTypeList, ListIndex: 0},
						Index{Type: IndexTypeMap, MapIndex: "certificate"},
						Index{Type: IndexTypeMap, MapIndex: "private_key_pem"},
					},
					"networking_poe_ssl_certs_0/name": []Index{
						Index{Type: IndexTypeMap, MapIndex: ".properties.networking_poe_ssl_certs"},
						Index{Type: IndexTypeMap, MapIndex: "value"},
						Index{Type: IndexTypeList, ListIndex: 0},
						Index{Type: IndexTypeMap, MapIndex: "name"},
					},
				}))
			})
		})
	})
})
