package people_test

import (
	"github.com/jtarchie/simulation/people"
	"image/color"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPeople(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "People Suite")
}

var _ = Describe("Person", func() {
	It("has dimensions", func() {
		person := &people.Person{}
		Expect(person.X()).To(BeEquivalentTo(0))
		Expect(person.Y()).To(BeEquivalentTo(0))

		person.SetXY(1,1)
		Expect(person.X()).To(BeEquivalentTo(1))
		Expect(person.Y()).To(BeEquivalentTo(1))
	})

	It("has a default color", func() {
		person := &people.Person{}
		Expect(person.Color()).To(Equal(color.RGBA{255,255,255,1}))
	})

	It("does not move", func(){
		person := &people.Person{}
		Expect(person.X()).To(BeEquivalentTo(0))
		Expect(person.Y()).To(BeEquivalentTo(0))

		person.Update(nil)
		Expect(person.X()).To(BeEquivalentTo(0))
		Expect(person.Y()).To(BeEquivalentTo(0))
	})
})