package fancyparser

import (
	"io/ioutil"
	"path"
	"strings"
)

// ExtractYAMLBytesInDir takes a path to a directory and returns a map where the
// keys are files in the directory and the values are the bytes of those files
func ExtractYAMLBytesInDir(pathToDir string) (map[string][]byte, error) {
	files, err := ioutil.ReadDir(pathToDir)
	if err != nil {
		// TODO: handle better
		return nil, nil // err
	}

	dirMap := make(map[string][]byte)

	for _, file := range files {
		filename := file.Name()
		if !strings.HasSuffix(filename, ".yml") {
			continue
		}

		fileBytes, err := ioutil.ReadFile(path.Join(pathToDir, file.Name()))
		if err != nil {
			return nil, err
		}
		dirMap[file.Name()] = fileBytes
	}

	return dirMap, nil
}
