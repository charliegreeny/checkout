package item

import (
	"errors"
	"fmt"

	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/charliegreeny/checkout/pkg/offer"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type getter struct {
	cache map[string]*Entity
	offerGetter model.IDGetter[*offer.Entity]
	db    *gorm.DB
	log    *zap.Logger
}

func NewGetter(db *gorm.DB, og model.IDGetter[*offer.Entity], l *zap.Logger) (model.IDGetter[*Entity], error) {
	g := &getter{
		cache:       map[string]*Entity{},
		offerGetter: og,
		db:          db,
		log: l,
	}
	if err := g.AddToCache(); err != nil {
		return nil, err
	}
	return g, nil
}

func (g getter) GetById(id string) (*Entity, error) {
	g.log.Info("added to cache", zap.Any("cache", g.cache))
	if e, ok := g.cache[id]; ok {
		return e, nil
	}
	return nil, model.ErrNotFound{Err: fmt.Errorf("item SKU %s not found", id)}
}

func (g getter) AddToCache() error {
	var items []*Entity
	if r := g.db.Find(&items); r.Error != nil {
		return r.Error
	}
	for _, i := range items {
		if i.OffersId != nil {
			o, err := g.offerGetter.GetById(*i.OffersId)
			if err != nil && !errors.As(err, &model.ErrNotFound{}){
				return err
			}
			i.Offer = o 
		}
		g.cache[i.SKU] = i
	}
	return nil
}
