package fancyparser_test

import (
	"io/ioutil"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/tile-config-generator/fancyparser"
)

var _ = Describe("Common", func() {
	Context("ExtractYMLBytesInDir", func() {
		var (
			pathToDir string
			dirBytes  map[string][]byte
			err       error
		)

		JustBeforeEach(func() {
			dirBytes, err = ExtractYAMLBytesInDir(pathToDir)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when the dir exists", func() {
			BeforeEach(func() {
				pathToDir, err = ioutil.TempDir("", "")
				Expect(err).ToNot(HaveOccurred())

				file1 := path.Join(pathToDir, "beep.yml")
				err = ioutil.WriteFile(file1, []byte("beep"), 0644)
				Expect(err).ToNot(HaveOccurred())

				file2 := path.Join(pathToDir, "boop.yml")
				err = ioutil.WriteFile(file2, []byte("boop"), 0644)
				Expect(err).ToNot(HaveOccurred())

				file3 := path.Join(pathToDir, "foo")
				err = ioutil.WriteFile(file3, []byte("foo"), 0644)
				Expect(err).ToNot(HaveOccurred())
			})

			It("reads in all the files with .yml suffix", func() {
				Expect(dirBytes).To(Equal(map[string][]byte{
					"beep.yml": []byte("beep"),
					"boop.yml": []byte("boop"),
				}))
			})
		})

		Context("when the dir doesn't exist", func() {
		})
	})
})
