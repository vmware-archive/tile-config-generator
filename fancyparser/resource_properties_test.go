package fancyparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("LookupResourceProperty", func() {
	var (
		resourceProperties []interface{}
		indexList          []Index
		value              interface{}
		err                error
	)

	JustBeforeEach(func() {
		value, err = LookupResourceProperty(indexList, resourceProperties)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("when the resource property is nested 2 indexes deep", func() {
		BeforeEach(func() {
			indexList = []Index{
				Index{Type: IndexTypeMap, MapIndex: "diego_database"},
				Index{Type: IndexTypeMap, MapIndex: "instances"},
			}
			resourceProperties = []interface{}{
				map[string]interface{}{
					"description":            "An datastore node for Diego",
					"identifier":             "diego_database",
					"instance_type_best_fit": "micro",
					"instance_type_id":       "medium",
					"instances":              3,
					"instances_best_fit":     3,
				},
			}
		})

		It("it looks up the second index", func() {
			Expect(value).To(Equal(3))
		})
	})

	Context("when the resource property is nested 3 indexes deep", func() {
		BeforeEach(func() {
			indexList = []Index{
				Index{Type: IndexTypeMap, MapIndex: "diego_database"},
				Index{Type: IndexTypeMap, MapIndex: "instance_type"},
				Index{Type: IndexTypeMap, MapIndex: "id"},
			}
			resourceProperties = []interface{}{
				map[string]interface{}{
					"description":            "An datastore node for Diego",
					"identifier":             "diego_database",
					"instance_type_best_fit": "micro",
					"instance_type_id":       "medium",
					"instances":              3,
					"instances_best_fit":     3,
				},
			}
		})

		It("regex matches the 2nd and 3rd index", func() {
			Expect(value).To(Equal("medium"))
		})
	})
})
