package entity

import (
	"errors"
	"time"
)


type Plain struct {
	Id string
	Name string
	Sku *Sku
	Price float64
	CreateAt time.Time
	UpdateAt time.Time
	ExpiredAt time.Time
}

func NewPlain(name string, price float64, sku *Sku, expiredAt time.Time) (*Plain, error) {
	newPlain := Plain {
		Name: name,
		Price: price,
		Sku: sku,
		ExpiredAt: expiredAt,
	}
	err := newPlain.validate()
	if (err != nil) {
		return nil, err
	}
	return &newPlain, nil
}

func (p *Plain) validate() error {
	if (p.Price <= 0) {
		return errors.New("price invalid")
	}
	if (!p.Sku.Active()) {
		return errors.New("sku inactivate")
	}
	expiredAtIsRetro := time.Now().After(p.ExpiredAt)
	if (expiredAtIsRetro) {
		return errors.New("expiredAt is invalid")
	}
	return nil
}

/*
- Um plano pode conter somente um SKU, ter um nome e um preço, data de validade

- O preço não pode ser negativo

- O plano não pode conter um SKU inativo

- A data de validade não pode ser menor que a data atual

*/