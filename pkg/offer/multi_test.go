package offer

import (
	"testing"
)

func Int(v int) *int {
	return &v
}

func Test_multiApplier_Apply(t *testing.T) {
	type args struct {
		quantityGot int
		price       int
	}
	tests := []struct {
		name  string
		offer *Entity
		args  args
		want  int
	}{
		{
			name: "perfectly matches multi-buy offer, returns amount from offer",
			offer: &Entity{
				ID:       "offer1",
				Quantity: Int(5),
				Price:    Int(250),
				Discount: nil,
			},
			args: args{
				quantityGot: 5, price: 100,
			},
			want: 250,
		},
		{
			name: "more than multi-buy offer, returns amount from offer plus normal price",
			offer: &Entity{
				ID:       "offer1",
				Quantity: Int(5),
				Price:    Int(250),
				Discount: nil,
			},
			args: args{
				quantityGot: 7, price: 100,
			},
			want: 450,
		},
		{
			name: "qualifiers for offer 3 times , returns amount from offer * 3",
			offer: &Entity{
				ID:       "offer1",
				Quantity: Int(5),
				Price:    Int(250),
				Discount: nil,
			},
			args: args{
				quantityGot: 15, price: 100,
			},
			want: 750,
		},
		{
			name: "does not qualifier, returns amount from args.QuantityGot * args.Price",
			offer: &Entity{
				ID:       "offer1",
				Quantity: Int(5),
				Price:    Int(250),
				Discount: nil,
			},
			args: args{
				quantityGot: 3, price: 100,
			},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := multiApplier{
				offer: tt.offer,
			}
			if got := m.Apply(tt.args.quantityGot, tt.args.price); got != tt.want {
				t.Errorf("multiApplier.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
