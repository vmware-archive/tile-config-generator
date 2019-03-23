package commands

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/pivotalservices/tile-config-generator/metadata"
	yaml "gopkg.in/yaml.v2"
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
			if err != nil {
				return nil, err
			}
			contents, err := ioutil.ReadAll(metadataFile)
			if err != nil {
				return nil, err
			}
			return contents, nil
		}
	}
	return nil, errors.New("no metadata file was found in provided .pivotal")
}

// PruneFilepath takes a filepath, substitutes any ~ for file paths and returns
// the absolute path
// TODO: test
func PruneFilepath(file string) (string, error) {
	if rune(file[0]) == os.PathSeparator {
		return file, nil
	} else if file[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return filepath.Join(usr.HomeDir, file[2:]), nil
	} else {
		absPath, err := filepath.Abs(file)
		if err != nil {
			return "", err
		}
		return absPath, err
	}
}

func ExtractBytes(pathToFile string) ([]byte, error) {
	// absPath, err := filepath.Abs(pathToFile)
	// if err != nil {
	// 	return nil, err
	// }

	// content, err := ioutil.ReadFile(absPath)
	content, err := ioutil.ReadFile(pathToFile)
	return content, err
}

func writeYamlFile(targetFile string, dataType interface{}) error {
	if dataType != nil {
		data, err := yaml.Marshal(dataType)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(targetFile, data, 0755)
	} else {
		return ioutil.WriteFile(targetFile, nil, 0755)
	}
}

func itemInSlice(item string, slice []string) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}

	return false
}
