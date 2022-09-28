package entity

import (
	"errors"
	"time"
)

type Sku struct {
	Id string
	Name string
	Description string
	CreateAt time.Time
	UpdateAt time.Time
	IsActive bool
	UnityMeasured string
	Quantity string
}

func NewSku(name string, unitMeasured string, quantity string) (*Sku, error) {
	newSku := Sku{
		Name: name,
		Description: name,
		CreateAt: time.Now(),
		UnityMeasured: quantity,
		IsActive: true,
		Quantity: quantity,	
	}
	err := newSku.validate()
	return &newSku, err
}

func (s *Sku) validate() error {
	if (s.Name == "") {
		return errors.New("name is empty")
	}
	return nil
}

func (s *Sku) CalculateConsumption(quantity float64, price float64) float64 {
	return quantity * price
}

func (s * Sku) Active() bool {
	return s.IsActive
}