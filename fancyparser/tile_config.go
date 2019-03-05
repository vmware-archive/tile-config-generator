package fancyparser

type ProductConfig struct {
	ProductName       string      `json:"product-name"`
	NetworkProperties interface{} `json:"network-properties"`
	ProductProperties interface{} `json:"product-properties"`
	ResourceConfig    interface{} `json:"resource-config"`
}

type OpsFiles struct {
}

type VarsFiles struct {
}
