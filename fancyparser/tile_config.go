package fancyparser

import (
	"io/ioutil"
	"path/filepath"
)

type ProductConfig struct {
	// found in product.yml
	ProductName       string      `json:"product-name"`
	NetworkProperties interface{} `json:"network-properties"`
	ProductProperties interface{} `json:"product-properties"`
	ResourceConfig    interface{} `json:"resource-config"`

	// found by parsing the ops files in features/ network/ optional/ & resource/ dirs
	NetworkOpsFiles  []OpsFile
	FeaturesOpsFiles []OpsFile
	OptionalOpsFiles []OpsFile
}

type TileConfigBytes struct {
	FeaturesOpsFiles   map[string][]byte
	NetworkOpsFiles    map[string][]byte
	OptionalOpsFiles   map[string][]byte
	ResourceOpsFiles   map[string][]byte
	ErrandVars         []byte
	Metadata           []byte
	ProductDefaultVars []byte
	Product            []byte
	ResourceVars       []byte
}

// TODO: test!!!
func GetTileConfigBytes(tileDir string) (TileConfigBytes, error) {
	t := TileConfigBytes{}
	featuresOps, err := ExtractYAMLBytesInDir(filepath.Join(tileDir, "features"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.FeaturesOpsFiles = featuresOps

	networksOps, err := ExtractYAMLBytesInDir(filepath.Join(tileDir, "network"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.NetworkOpsFiles = networksOps

	optionalOps, err := ExtractYAMLBytesInDir(filepath.Join(tileDir, "optional"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.OptionalOpsFiles = optionalOps

	resourceOps, err := ExtractYAMLBytesInDir(filepath.Join(tileDir, "resource"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.ResourceOpsFiles = resourceOps

	errandVars, err := ioutil.ReadFile(filepath.Join(tileDir, "errand-vars.yml"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.ErrandVars = errandVars

	metadata, err := ioutil.ReadFile(filepath.Join(tileDir, "metadata.yml"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.Metadata = metadata

	productDefaultVars, err := ioutil.ReadFile(filepath.Join(tileDir, "product-default-vars.yml"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.ProductDefaultVars = productDefaultVars

	product, err := ioutil.ReadFile(filepath.Join(tileDir, "product.yml"))
	if err != nil {
		return TileConfigBytes{}, err
	}
	t.Product = product

	resourceVars, err := ioutil.ReadFile(filepath.Join(tileDir, "resource-vars.yml"))
	if err != nil {
		return TileConfigBytes{}, err
	}

	t.ResourceVars = resourceVars
	return t, nil
}
