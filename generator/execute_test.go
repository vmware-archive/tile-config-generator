package generator_test

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotalservices/tile-config-generator/generator"
)

type Template struct {
	NetworkProperties interface{} `yaml:"network-properties"`
	ProductProperties interface{} `yaml:"product-properties"`
	ResourceConfig    interface{} `yaml:"resource-config,omitempty"`
	ErrandConfig      interface{} `yaml:"errand-config,omitempty"`
}

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
			testGen        = path.Join(pwd, "_testGen")
			tmpPath        = path.Join(testGen, "templates")
			tempDir string = os.TempDir()
			zipPath string
		)
		BeforeEach(func() {

		})
		AfterEach(func() {
			err := os.RemoveAll(testGen)
			Expect(err).ShouldNot(HaveOccurred())
			os.Remove(zipPath)
		})

		It("Should generate files for p-healthwatch", func() {
			zipPath = path.Join(tempDir, "p-healthwatch.pivotal")
			err := createZipFile("fixtures/p_healthwatch.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should generate files for pas", func() {
			zipPath = path.Join(tempDir, "pas.pivotal")
			err := createZipFile("fixtures/pas.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for pas 2.2", func() {
			zipPath = path.Join(tempDir, "pas.pivotal")
			err := createZipFile("fixtures/pas_2_2.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for mysql_v2", func() {
			zipPath = path.Join(tempDir, "mysql.pivotal")
			err := createZipFile("fixtures/mysql_v2.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for scs", func() {
			zipPath = path.Join(tempDir, "scs.pivotal")
			err := createZipFile("fixtures/scs.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for srt", func() {
			zipPath = path.Join(tempDir, "srt.pivotal")
			err := createZipFile("fixtures/srt.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, true, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should generate files for push notifications", func() {
			zipPath = path.Join(tempDir, "push_notifications.pivotal")
			err := createZipFile("fixtures/p_push_notifications.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should generate files for pivotal cloud cache", func() {
			zipPath = path.Join(tempDir, "pcc.pivotal")
			err := createZipFile("fixtures/cloudcache.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for rabbitmq", func() {
			zipPath = path.Join(tempDir, "rabbit.pivotal")
			err := createZipFile("fixtures/rabbit-mq.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for redis", func() {
			zipPath = path.Join(tempDir, "redis.pivotal")
			err := createZipFile("fixtures/p-redis.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for apigee", func() {
			zipPath = path.Join(tempDir, "apigee.pivotal")
			err := createZipFile("fixtures/apigee.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should generate files for pks", func() {
			zipPath = path.Join(tempDir, "pks.pivotal")
			err := createZipFile("fixtures/pks.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
			template, err := unmarshalProduct(path.Join(tmpPath, "pivotal-container-service", "1.1.3-build.11", "product.yml"))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template.NetworkProperties).ShouldNot(BeNil())
			Expect(template.ResourceConfig).ShouldNot(BeNil())
		})
		It("Should generate files for nsx-t", func() {
			zipPath = path.Join(tempDir, "nsx-t.pivotal")
			err := createZipFile("fixtures/nsx-t.yml", zipPath)
			Expect(err).ShouldNot(HaveOccurred())
			gen = generator.NewExecutor(zipPath, tmpPath, false, true)
			err = gen.Generate()
			Expect(err).ShouldNot(HaveOccurred())
			template, err := unmarshalProduct(path.Join(tmpPath, "VMware-NSX-T", "2.2.1.9149087", "product.yml"))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(template.NetworkProperties).Should(BeNil())
			Expect(template.ResourceConfig).Should(BeNil())
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

func unmarshalProduct(targetFile string) (*Template, error) {
	template := &Template{}
	yamlFile, err := ioutil.ReadFile(targetFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, template)
	if err != nil {
		return nil, err
	}
	return template, nil
}
