package fancyparser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFancyparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fancyparser Suite")
}
