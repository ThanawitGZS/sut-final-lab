package unit

import (
	"github.com/ThanawitGZS/sut-final-lab/entity"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"testing"
)

func TestProductNegativePrice(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`price is invalid`,func (t *testing.T)  {
		p := entity.Products {
			Name:"A",
			Price:10111,
			SKU:"ABC00001",
		}

		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo((BeTrue()))
		g.Expect(err).NotTo((BeNil()))
		g.Expect(err.Error()).To(Equal("Price must be between 1 and 1000"))
	})
}