package generator_test

import (
	"io/ioutil"

	"github.com/calebwashburn/tile-config-template-generator/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JobType", func() {
	Context("HasPersistentDisk", func() {
		It("Should have persistent disk", func() {
			fileData, err := ioutil.ReadFile("fixtures/pas.yml")
			Expect(err).ShouldNot(HaveOccurred())
			metadata, err := generator.NewMetadata(fileData)
			Expect(err).ShouldNot(HaveOccurred())
			job, err := metadata.GetJob("mysql")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(job.HasPersistentDisk()).Should(BeTrue())
		})
		It("Should not have persistent disk", func() {
			fileData, err := ioutil.ReadFile("fixtures/pas.yml")
			Expect(err).ShouldNot(HaveOccurred())
			metadata, err := generator.NewMetadata(fileData)
			Expect(err).ShouldNot(HaveOccurred())
			job, err := metadata.GetJob("router")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(job.HasPersistentDisk()).Should(BeFalse())
		})
	})
})
