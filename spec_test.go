package swaggen_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/goline/swaggen"
)

var _ = Describe("Spec", func() {
	It("New should panic if version is not supported", func(d Done) {
		versions := []string{"2.1", "2"}
		c := make(chan bool, len(versions))
		for _, version := range versions {
			go func(version string, c chan bool) {
				defer GinkgoRecover()
				defer func() {
					r := recover()
					Expect(r).NotTo(BeNil())
					Expect(r).To(Equal(swaggen.ErrInvalidVersion))
					c <- true
				}()

				swaggen.New(version)
			}(version, c)
		}

		for i := 0; i < len(versions); i++ {
			Expect(<-c).To(BeTrue())
		}

		close(d)
	})
})
