package cart

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_service_GetById(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Entity
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				db: tt.fields.db,
			}
			got, err := s.GetById(tt.args.id)
			tt.wantErr(t, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}
