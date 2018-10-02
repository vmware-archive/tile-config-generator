package commands

type PivnetConfiguration struct {
	Token   string `long:"token" env:"PIVNET_TOKEN" description:"Pivnet Token"`
	Slug    string `long:"product-slug" env:"PIVNET_PRODUCT_SLUG" description:"Pivnet Product Slug"`
	Version string `long:"product-version" env:"PIVNET_PRODUCT_VERSION" description:"Pivnet Product Version"`
	Glob    string `long:"product-glob" env:"PIVNET_PRODUCT_GLOB" description:"Pivnet Product Glob" default:"*.pivotal"`
}
