package aplication_test

import (
	"testing"

	"github.com/josenunesti/baas/aplication"
	"github.com/stretchr/testify/assert"
)

func TestDeveRetornarUmPrecoDeUmSkuValido(t *testing.T) {

	skuRepositoryMemory := infra_respository.SkuRepositoryInMemory()
	skuRepositoryMemory := infra_respository.SkuRepositoryInMemory()
	useCase := aplication.NewUseCaseGetSku(skuRepositoryMemory)
	skuId := "123456"
	result, err := useCase.Exec(skuId)
	assert.NotNil(err)
	expectedPrice := 0.005
	assert.Equal(t, expectedPrice, result.Price)
}