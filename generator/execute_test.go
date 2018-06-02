package generator_test

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/calebwashburn/tile-config-template-generator/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Executor", func() {
	Context("Generate", func() {
		var (
			gen              *generator.Executor
			pwd, _           = os.Getwd()
			tmpDirName       = "_testGen"
			tmpPath          = path.Join(pwd, tmpDirName)
			controlOutputDir string
			fileData         []byte
		)
		BeforeEach(func() {
			gen = &generator.Executor{}
			os.MkdirAll(tmpPath, 0700)
			controlOutputDir, _ = ioutil.TempDir(tmpPath, "templates")
			var err error
			fileData, err = ioutil.ReadFile("fixtures/p_healthwatch.yml")
			Expect(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			err := os.RemoveAll(tmpPath)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should create output template with network properties", func() {
			template, err := gen.Generate(fileData)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
			Expect(template.NetworkProperties).ShouldNot(BeNil())
		})
		XIt("Should create output template with product properties", func() {
			template, err := gen.Generate(fileData)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
			Expect(template.ProductProperties).ShouldNot(BeNil())
		})
		It("Should create output template with resource config properties", func() {
			template, err := gen.Generate(fileData)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
			Expect(template.ResourceConfig).ShouldNot(BeNil())
		})
	})
})
