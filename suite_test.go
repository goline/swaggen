package swaggen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSwaggen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Swaggen Suite")
}
