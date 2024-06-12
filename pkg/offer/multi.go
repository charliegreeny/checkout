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
		timesQual := quantityGot / *m.offer.Quantity
		total = *m.offer.Price * timesQual
		if remainder := quantityGot % *m.offer.Quantity; remainder > 0 {
			a := remainder * price
			total += a
			return total
		}
		return total
	}
	return quantityGot * price
}
