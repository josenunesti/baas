package entity_test

import (
	"testing"

	"github.com/josenunesti/baas/domain/entity"
	"github.com/stretchr/testify/assert"
)


func TestDeveCriarUmSkuValido(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, error := entity.NewSku(name, unityMeasured, quantity)
	assert.Nil(t, error)
	assert.Equal(t, sku.Name, name)
}

func TestNaoDeveCriarUmSkuQuantoONomeEVazio(t *testing.T) {
	name := ""
	unityMeasured := "GiB"
	quantity := "second"
	sku, error := entity.NewSku(name, unityMeasured, quantity)
	assert.Equal(t, error.Error(), "name is empty")
	assert.Equal(t, sku.Name, "")
}

func TestDeveCalcularConsumoDeUmSku(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, err := entity.NewSku(name, unityMeasured, quantity)
	if (err != nil) {
		t.Error(err)
	}
	quantityConsumption := 60.0
	price := float64(0.001)
	totalAmount := sku.CalculateConsumption(quantityConsumption, price)
	totalAmountExpected := float64(0.06)
	assert.Equal(t, totalAmountExpected, totalAmount)
}

func TestDeveCalcularConsumoDeUmSkuComQuantidadeMenor(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, err := entity.NewSku(name, unityMeasured, quantity)
	if (err != nil) {
		t.Error(err)
	}
	quantityConsumption := 0.5
	price := float64(0.001)
	totalAmount := sku.CalculateConsumption(quantityConsumption, price)
	totalAmountExpected := float64(0.0005)
	assert.Equal(t, totalAmountExpected, totalAmount)
}

func TestDeveRetornarTrueQuantoSkuEstaAtivo(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, err := entity.NewSku(name, unityMeasured, quantity)
	if (err != nil) {
		t.Error(err)
	}
	isActive := sku.Active()
	assert.True(t, isActive)
}

func TestDeveRetornarFalseQuantoSkuEstaDesativado(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, err := entity.NewSku(name, unityMeasured, quantity)
	if (err != nil) {
		t.Error(err)
	}
	sku.IsActive = false
	isActive := sku.Active()
	assert.False(t, isActive)
}
