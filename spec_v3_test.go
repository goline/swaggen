package swaggen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/goline/swaggen"
)

var _ = Describe("SpecV3", func() {
	It("should allow to scan file's comments", func() {
		s := swaggen.New("3")
		err := s.Scan(swaggen.FileSource{Path: "./fixtures/doc.go"})
		Expect(err).To(BeNil())

		s3 := s.(*swaggen.SpecV3)
		Expect(s3.Info.Title).To(Equal("Identity API"))
		Expect(s3.Info.Contact.Name).To(Equal("Long Do"))
		Expect(s3.Info.Contact.Email).To(Equal("me@dotronglong.com"))
		Expect(s3.Info.License.Name).To(Equal("MIT"))
		Expect(s3.Info.Version).To(Equal("1.0.1"))
	})
})
