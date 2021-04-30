package awesomeProject

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("FizzBuzz", func() {
	ginkgo.Describe("FizzBuzz", func() {
		ginkgo.It("should return string given int", func() {
			got := FizzBuzz(1)
			gomega.Expect(got).To(gomega.Equal("1"))
		})
	})
})
