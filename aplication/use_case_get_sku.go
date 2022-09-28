package aplication

import (
	"errors"
	"fmt"

	"github.com/josenunesti/baas/domain/entity"
)


type DtoOutputUseCaseGetSku struct {
	NameSku string
	Price float64
}

type SkuRepository interface {
	GetSkuById(skuId string) (entity.Sku, error) 
}

type PlainRepository interface {
	GetPlainBySku(skuId string) (entity.Plain, error) 
}

type UseCaseGetSku struct {
	skuRepository SkuRepository
	plainRepository PlainRepository
}

func NewUseCaseGetSku(skuRepository SkuRepository, plainRepository PlainRepository) *UseCaseGetSku {
	return &UseCaseGetSku{
		skuRepository: skuRepository,
		plainRepository: plainRepository,
	}
}

func (u *UseCaseGetSku) Exec(skuId string) (DtoOutputUseCaseGetSku, error) {
	dtoOutput := DtoOutputUseCaseGetSku{}
	sku, err := u.skuRepository.GetSkuById(skuId)
	if (err != nil) {
		return dtoOutput, err
	}
	if (!sku.Active()) {
		msgError := fmt.Sprintf("sku: %s is disabled", skuId)
		return dtoOutput, errors.New(msgError)
	}
	plain, err := u.plainRepository.GetPlainBySku(skuId)
	if (err != nil) {
		return dtoOutput, err
	}
	dtoOutput.NameSku = sku.Name
	dtoOutput.Price = plain.Price
	return dtoOutput, nil
}
