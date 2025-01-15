package unit

import (
	"testing"

	"github.com/ThanawitGZS/sut-final-lab/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestProductNegativeSKU(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`sku is invalid`, func(t *testing.T) {
		p := entity.Products{
			Name:  "A",
			Price: 101,
			SKU:   "01245ADC",
		}

		ok, err := govalidator.ValidateStruct(p)
		g.Expect(ok).NotTo((BeTrue()))
		g.Expect(err).NotTo((BeNil()))
		g.Expect(err.Error()).To(Equal("SKU must start with 3 uppercase English letters (A-Z) followed by 5 digits (0-9)"))
	})
}
