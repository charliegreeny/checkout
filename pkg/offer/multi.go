package offer

type multiApplier struct {
	offer *Entity
}

func newMultiApplier(o *Entity) Applier {
	return &multiApplier{o}
}

func (m multiApplier) Apply(quantityGot, price int) int {
	total := 0
	if quantityGot >= *m.offer.Quantity {
		timesQual := *m.offer.Quantity / quantityGot

		total = *m.offer.Price * timesQual
		if remainder := *m.offer.Quantity % quantityGot; remainder > 0 {
			a := remainder * price
			total += a
			return total
		}
		return total
	}
	return quantityGot * price
}
