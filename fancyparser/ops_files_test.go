package fancyparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("Ops Files", func() {
	Context("GetOpsFileMapFromDirBytes", func() {
		var (
			dirBytes   map[string][]byte
			opsFileMap map[string]OpsFile
			err        error
		)

		JustBeforeEach(func() {
			opsFileMap, err = GetOpsFileMapFromDirBytes(dirBytes)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when all the unmarshalling succeeds", func() {
			BeforeEach(func() {
				dirBytes = map[string][]byte{
					"uaa_database-external.yml": []byte(`- type: replace
  path: /product-properties/.properties.uaa_database?
  value:
    value: external
- type: replace
  path: /product-properties/.properties.uaa_database.external.host?
  value:
    value: ((uaa_database/external/host))
- type: replace
  path: /product-properties/.properties.uaa_database.external.port?
  value:
    value: ((uaa_database/external/port))
`),
					"uaa_ldap.yml": []byte(`- type: replace
  path: /product-properties/.properties.uaa?
  value:
    value: ldap
- type: remove
  path: /product-properties/.properties.uaa.internal.password_min_length?
`),
				}

			})

			It("unmarshals the ops bytes into a opsfile map", func() {
				Expect(opsFileMap).To(Equal(map[string]OpsFile{
					"uaa_database-external.yml": OpsFile{
						Ops: []Operation{
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
						},
					},
					"uaa_ldap.yml": OpsFile{
						Ops: []Operation{
							Operation{
								Path: "/product-properties/.properties.uaa?",
								Type: OperationTypeReplace,
								Value: map[string]interface{}{
									"value": "ldap",
								},
							},
							Operation{
								Path: "/product-properties/.properties.uaa.internal.password_min_length?",
								Type: OperationTypeRemove,
							},
						},
					},
				}))
			})
		})
	})
})
