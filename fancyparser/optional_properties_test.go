package fancyparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("Optional Properties", func() {
	var (
		productProperties interface{}
		opsFile           OpsFile
		err               error
	)

	BeforeEach(func() {
		opsFile = OpsFile{}
	})

	Context("CheckOptionalIncludeAndGetIndexMap", func() {
		JustBeforeEach(func() {
			err = opsFile.CheckOptionalIncludeAndGetIndexMap(productProperties)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the property is a collection", func() {
			Context("when the collection lengths match", func() {
				BeforeEach(func() {
					opsFile.Ops = []Operation{
						Operation{
							Path: "/product-properties/.properties.push_apps_manager_footer_links?",
							Type: OperationTypeReplace,
							Value: map[string]interface{}{
								"value": []interface{}{
									map[string]interface{}{
										"href": "((push_apps_manager_footer_links_0/href))",
										"name": "((push_apps_manager_footer_links_0/name))",
									},
									map[string]interface{}{
										"href": "((push_apps_manager_footer_links_1/href))",
										"name": "((push_apps_manager_footer_links_1/name))",
									},
								},
							},
						},
					}

					productProperties = map[string]interface{}{
						".properties.push_apps_manager_footer_links": map[string]interface{}{
							"configurable": true,
							"credential":   false,
							"optional":     true,
							"type":         "collection",
							"value": []interface{}{
								map[string]interface{}{
									"href": map[string]interface{}{
										"value": "google.com",
									},
									"guid": map[string]interface{}{
										"value": "beep-boop",
									},
								},
								map[string]interface{}{
									"href": map[string]interface{}{
										"value": "yahoo.com",
									},
									"guid": map[string]interface{}{
										"value": "foo-bar",
									},
								},
							},
						},
					}
				})

				It("includes the ops file and collects the index maps", func() {
					Expect(opsFile.Include).To(BeTrue())
					Expect(opsFile.IndexMap).To(BeEquivalentTo(IndexMap{
						"push_apps_manager_footer_links_0/href": []Index{
							Index{Type: IndexTypeMap, MapIndex: ".properties.push_apps_manager_footer_links"},
							Index{Type: IndexTypeMap, MapIndex: "value"},
							Index{Type: IndexTypeList, ListIndex: 0},
							Index{Type: IndexTypeMap, MapIndex: "href"},
						},
						"push_apps_manager_footer_links_0/name": []Index{
							Index{Type: IndexTypeMap, MapIndex: ".properties.push_apps_manager_footer_links"},
							Index{Type: IndexTypeMap, MapIndex: "value"},
							Index{Type: IndexTypeList, ListIndex: 0},
							Index{Type: IndexTypeMap, MapIndex: "name"},
						},
						"push_apps_manager_footer_links_1/href": []Index{
							Index{Type: IndexTypeMap, MapIndex: ".properties.push_apps_manager_footer_links"},
							Index{Type: IndexTypeMap, MapIndex: "value"},
							Index{Type: IndexTypeList, ListIndex: 1},
							Index{Type: IndexTypeMap, MapIndex: "href"},
						},
						"push_apps_manager_footer_links_1/name": []Index{
							Index{Type: IndexTypeMap, MapIndex: ".properties.push_apps_manager_footer_links"},
							Index{Type: IndexTypeMap, MapIndex: "value"},
							Index{Type: IndexTypeList, ListIndex: 1},
							Index{Type: IndexTypeMap, MapIndex: "name"},
						},
					}))
				})
			})

			Context("when the actual collection length is 0", func() {
				BeforeEach(func() {
					opsFile.Ops = []Operation{
						Operation{
							Path: "/product-properties/.properties.push_apps_manager_footer_links?",
							Type: OperationTypeReplace,
							Value: map[string]interface{}{
								"value": []interface{}{
									map[string]interface{}{
										"href": "((push_apps_manager_footer_links_0/href))",
										"name": "((push_apps_manager_footer_links_0/name))",
									},
									map[string]interface{}{
										"href": "((push_apps_manager_footer_links_1/href))",
										"name": "((push_apps_manager_footer_links_1/name))",
									},
								},
							},
						},
					}

					productProperties = map[string]interface{}{
						".properties.push_apps_manager_footer_links": map[string]interface{}{
							"configurable": true,
							"credential":   false,
							"optional":     true,
							"type":         "collection",
							"value":        nil,
						},
					}
				})

				It("doesn't include the ops file", func() {
					Expect(opsFile.Include).To(BeFalse())
				})
			})

			Context("when the collection lengths don't match", func() {
				BeforeEach(func() {
					opsFile.Ops = []Operation{
						Operation{
							Path: "/product-properties/.properties.push_apps_manager_footer_links?",
							Type: OperationTypeReplace,
							Value: map[string]interface{}{
								"value": []interface{}{
									map[string]interface{}{
										"href": "((push_apps_manager_footer_links_0/href))",
										"name": "((push_apps_manager_footer_links_0/name))",
									},
									map[string]interface{}{
										"href": "((push_apps_manager_footer_links_1/href))",
										"name": "((push_apps_manager_footer_links_1/name))",
									},
								},
							},
						},
					}

					productProperties = map[string]interface{}{
						".properties.push_apps_manager_footer_links": map[string]interface{}{
							"configurable": true,
							"credential":   false,
							"optional":     true,
							"type":         "collection",
							"value": []interface{}{
								map[string]interface{}{
									"href": map[string]interface{}{
										"value": "google.com",
									},
									"guid": map[string]interface{}{
										"value": "beep-boop",
									},
								},
							},
						},
					}
				})

				It("doesn't include the ops file", func() {
					Expect(opsFile.Include).To(BeFalse())
				})
			})
		})

		Context("when the property isn't a collecton", func() {
			Context("when the ops file is excluded", func() {
				BeforeEach(func() {
					opsFile.Ops = []Operation{
						Operation{
							Path: "/product-properties/.uaa.service_provider_key_password?",
							Type: OperationTypeReplace,
							Value: map[string]interface{}{
								"value": map[string]interface{}{
									"secret": "((uaa/service_provider_key_password))",
									"beep":   "((boop))",
								},
							},
						},
					}

					productProperties = map[string]interface{}{
						".uaa.service_provider_key_password": map[string]interface{}{
							"configurable": true,
							"credential":   true,
							"optional":     true,
							"type":         "secret",
							"value": map[string]interface{}{
								"secret": "***",
							},
						},
					}
				})

				It("excludes the ops file", func() {
					Expect(opsFile.Include).To(BeFalse())
				})
			})

			Context("when the ops file is included", func() {
				BeforeEach(func() {
					opsFile.Ops = []Operation{
						Operation{
							Path: "/product-properties/.uaa.service_provider_key_password?",
							Type: OperationTypeReplace,
							Value: map[string]interface{}{
								"value": map[string]interface{}{
									"secret": "((uaa/service_provider_key_password))",
								},
							},
						},
					}

					productProperties = map[string]interface{}{
						".uaa.service_provider_key_password": map[string]interface{}{
							"configurable": true,
							"credential":   true,
							"optional":     true,
							"type":         "secret",
							"value": map[string]interface{}{
								"secret": "***",
							},
						},
					}
				})

				It("includes the ops file and collects the index maps", func() {
					Expect(opsFile.Include).To(BeTrue())
					Expect(opsFile.IndexMap).To(Equal(IndexMap{
						"uaa/service_provider_key_password": []Index{
							Index{Type: IndexTypeMap, MapIndex: ".uaa.service_provider_key_password"},
							Index{Type: IndexTypeMap, MapIndex: "value"},
							Index{Type: IndexTypeMap, MapIndex: "secret"},
						},
					}))
				})
			})

			// TODO: test!!
			Context("when the ops file is excluded", func() {
			})
		})
	})
})
