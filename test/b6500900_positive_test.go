package unit

import (
	"github.com/ThanawitGZS/sut-final-lab/entity"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"testing"
)

func TestProductPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`product is valid`,func (t *testing.T)  {
		p := entity.Products {
			Name:"A",
			Price:12.00,
			SKU:"ABC00001",
		}

		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}