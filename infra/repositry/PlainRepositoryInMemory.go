package repository

import (
	"github.com/josenunesti/baas/domain/entity"
	"github.com/josenunesti/baas/entity"
)

type PlainRepositoryInMemory struct {
	Skus []entity.Sku
}

func (s *PlainRepositoryInMemory) GetSkuById(skuId string) (*entity.Sku, error) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	skuFake, _ := entity.NewSku(name, unityMeasured, quantity)
	return skuFake, nil
}