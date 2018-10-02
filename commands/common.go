package commands

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"regexp"
)

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
