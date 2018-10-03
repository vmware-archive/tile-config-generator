package commands

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"regexp"

	"github.com/pivotalservices/tile-config-generator/metadata"
)

func getProvider(pathToPivotalFile string, pivnet *PivnetConfiguration) (metadata.Provider, error) {
	if pathToPivotalFile != "" {
		return metadata.NewFileProvider(pathToPivotalFile), nil
	} else {
		if pivnet.Token == "" || pivnet.Slug == "" || pivnet.Version == "" {
			return nil, errors.New("Must provide either --pivotal-file-path or pivnet --token --product-slug --product-version")
		}
		return metadata.NewPivnetProvider(pivnet.Token, pivnet.Slug, pivnet.Version, pivnet.Glob), nil
	}
}
func extractMetadataBytes(pathToPivotalFile string) ([]byte, error) {
	zipReader, err := zip.OpenReader(pathToPivotalFile)
	if err != nil {
		return nil, err
	}

	defer zipReader.Close()

	for _, file := range zipReader.File {
		metadataRegexp := regexp.MustCompile("metadata/.*\\.yml")
		matched := metadataRegexp.MatchString(file.Name)

		if matched {
			metadataFile, err := file.Open()
			contents, err := ioutil.ReadAll(metadataFile)
			if err != nil {
				return nil, err
			}
			return contents, nil
		}
	}
	return nil, errors.New("no metadata file was found in provided .pivotal")
}
