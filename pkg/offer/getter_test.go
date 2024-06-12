package offer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApplier(t *testing.T) {
	validE := &Entity{
		ID:       "offer1",
		Quantity: Int(1),
		Price:    Int(1),
		Discount: nil,
	}
	invalidE := &Entity{
		ID:       "offer1",
		Quantity: nil,
		Price:    nil,
		Discount: nil,
	}
	tests := []struct {
		name    string
		args    *Entity
		want    Applier
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "valid offer to apply, returns ",
			args:    validE,
			want:    &multiApplier{offer: validE},
			wantErr: assert.NoError,
		},
		{
			name:   "invalid offer to apply, returns ",
			args:    invalidE,
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetApplier(tt.args)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
