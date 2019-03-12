package fancyparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("Features Properties", func() {
	var (
		productProperties interface{}
		productConfig     interface{}
		opsFile           OpsFile
		err               error
	)

	Context("GetPropertyNameFromPath", func() {
		It("Converts the path string to a property name", func() {
			path := "/product-properties/.properties.uaa.saml.sso_url?"
			propertyName := GetPropertyNameFromPath(path)
			Expect(propertyName).To(Equal(".properties.uaa.saml.sso_url"))
		})
	})

	Context("CheckFeatureIncludeAndGetIndexMap", func() {
		JustBeforeEach(func() {
			err = opsFile.CheckFeatureIncludeAndGetIndexMap(productProperties, productConfig)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the ops file is included", func() {
			BeforeEach(func() {
				opsFile.Ops = []Operation{
					Operation{
						Path: "/product-properties/.properties.uaa_database?",
						Type: OperationTypeReplace,
						Value: map[string]interface{}{
							"value": "external",
						},
					},
					Operation{
						Path: "/product-properties/.properties.uaa_database.external.host?",
						Type: OperationTypeReplace,
						Value: map[string]interface{}{
							"value": "((uaa_database/external/host))",
						},
					},
					Operation{
						Path: "/product-properties/.properties.uaa_database.external.port?",
						Type: OperationTypeReplace,
						Value: map[string]interface{}{
							"value": "((uaa_database/external/port))",
						},
					},
				}

				productProperties = map[string]interface{}{
					".properties.uaa_database": map[string]interface{}{
						"configurable":    true,
						"credential":      false,
						"optional":        false,
						"selected_option": "external",
						"type":            "selector",
						"value":           "external",
					},
				}

				productConfig = map[string]interface{}{
					".properties.uaa_database": map[string]interface{}{
						"value": "internal_mysql",
					},
				}
			})

			It("includes the ops file and collects the index maps", func() {
				Expect(opsFile.Include).To(BeTrue())
				Expect(opsFile.Exclude).To(BeFalse())
				Expect(opsFile.IndexMap).To(Equal(IndexMap{
					"uaa_database/external/host": []Index{
						Index{Type: IndexTypeMap, MapIndex: ".properties.uaa_database.external.host"},
						Index{Type: IndexTypeMap, MapIndex: "value"},
					},
					"uaa_database/external/port": []Index{
						Index{Type: IndexTypeMap, MapIndex: ".properties.uaa_database.external.port"},
						Index{Type: IndexTypeMap, MapIndex: "value"},
					},
				}))
			})
		})

		Context("when the ops file is excluded", func() {
			BeforeEach(func() {
				opsFile.Ops = []Operation{
					Operation{
						Path:  "/product-properties/.properties.uaa_database?",
						Type:  OperationTypeReplace,
						Value: "internal",
					},
					Operation{
						Path:  "/product-properties/.properties.uaa_database.external.host?",
						Type:  OperationTypeReplace,
						Value: "((uaa_database/external/host))",
					},
					Operation{
						Path:  "/product-properties/.properties.uaa_database.external.port?",
						Type:  OperationTypeReplace,
						Value: "((uaa_database/external/port))",
					},
				}
			})
		})
	})
})
