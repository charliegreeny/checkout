package offer

import (
	"errors"

	"github.com/charliegreeny/checkout/pkg/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type getter struct {
	db *gorm.DB
}

func NewGetter(db *gorm.DB, l *zap.Logger) model.IDGetter[*Entity] {
	return getter{
		db:  db,
	}
}

func (g getter) GetById(id string) (*Entity, error) {
	var e *Entity
	if r := g.db.First(&e, "id = ?", id); r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return nil, model.ErrNotFound{Err: r.Error}
		}
		return nil, r.Error
	}
	return e, nil
}

func GetApplier(e *Entity) (Applier, error) {
	if e.Quantity != nil && e.Price != nil {
		return newMultiApplier(e), nil
	}
	return nil, errors.New("no offer to apply")
}
