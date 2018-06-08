package generator_test

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"path"

	"github.com/calebwashburn/tile-config-template-generator/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Executor", func() {
	Context("CreateTemplate", func() {
		var (
			gen      *generator.Executor
			metadata *generator.Metadata
		)
		BeforeEach(func() {
			gen = &generator.Executor{}
			fileData, err := ioutil.ReadFile("fixtures/p_healthwatch.yml")
			Expect(err).ShouldNot(HaveOccurred())
			metadata, err = generator.NewMetadata(fileData)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should create output template with network properties", func() {
			template, err := gen.CreateTemplate(metadata)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
			Expect(template.NetworkProperties).ShouldNot(BeNil())
		})
		It("Should create output template with product properties", func() {
			template, err := gen.CreateTemplate(metadata)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
			Expect(template.ProductProperties).ShouldNot(BeNil())
		})
		It("Should create output template with resource config properties", func() {
			template, err := gen.CreateTemplate(metadata)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template).ShouldNot(BeNil())
			Expect(template.ResourceConfig).ShouldNot(BeNil())
		})
	})

	Context("Generate", func() {
		var (
			gen     *generator.Executor
			pwd, _         = os.Getwd()
			tmpPath        = path.Join(pwd, "_testGen", "templates")
			tempDir string = os.TempDir()
			zipPath string
		)
		BeforeEach(func() {

		})
		AfterEach(func() {
			// err := os.RemoveAll(tmpPath)
			// Expect(err).ShouldNot(HaveOccurred())
			//
			// err = os.Remove(filePath)
			// Expect(err).ShouldNot(HaveOccurred())

			os.Remove(zipPath)
		})

		It("Should generate files for p-healthwatch", func() {
			zipPath = path.Join(tempDir, "p-healthwatch.pivotal")
			err := createZipFile("fixtures/p_healthwatch.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should generate files for pas", func() {
			zipPath = path.Join(tempDir, "pas.pivotal")
			err := createZipFile("fixtures/pas.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})

func createZipFile(metadataFile string, targetFile string) error {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	fileData, err := ioutil.ReadFile(metadataFile)
	if err != nil {
		return err
	}
	f, err := w.Create("metadata/metadata.yml")
	if err != nil {
		return err
	}
	_, err = f.Write(fileData)
	if err != nil {
		return err
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(targetFile, buf.Bytes(), 0755)
}
