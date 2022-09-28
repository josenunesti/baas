package entity_test

import (
	"testing"
	"time"

	"github.com/josenunesti/baas/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestDeveCriarUmPlanoValido(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, _ := entity.NewSku(name, unityMeasured, quantity)

	namePlain := "planoA"
	price := 0.0004
	expiredAt := time.Now().AddDate(1, 1, 0)
	plain, err := entity.NewPlain(namePlain, price, sku, expiredAt)
	assert.Nil(t, err)
	assert.Equal(t, namePlain, plain.Name)
}

func TestDeveErroQuandoDataDeExpiracaoRetroativa(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, _ := entity.NewSku(name, unityMeasured, quantity)

	namePlain := "planoA"
	price := 0.0004
	expiredAt := time.Now().AddDate(-1, 1, 0)
	plain, err := entity.NewPlain(namePlain, price, sku, expiredAt)
	assert.NotNil(t, err)
	expectedError := "expiredAt is invalid"
	assert.Equal(t, expectedError, err.Error())
	assert.Nil(t, plain)
}

func TestDeveErroQuandoSkuEstaInativo(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, _ := entity.NewSku(name, unityMeasured, quantity)
	sku.IsActive = false

	namePlain := "planoA"
	price := 0.0004
	expiredAt := time.Now().AddDate(-1, 1, 0)
	plain, err := entity.NewPlain(namePlain, price, sku, expiredAt)
	assert.NotNil(t, err)
	expectedError := "sku inactivate"
	assert.Equal(t, expectedError, err.Error())
	assert.Nil(t, plain)
}


func TestDeveErroQuandoPrecoEstaZerado(t *testing.T) {
	name := "rabbit-mq"
	unityMeasured := "GiB"
	quantity := "second"
	sku, _ := entity.NewSku(name, unityMeasured, quantity)

	namePlain := "planoA"
	price := 0.0
	expiredAt := time.Now().AddDate(1, 1, 0)
	plain, err := entity.NewPlain(namePlain, price, sku, expiredAt)
	assert.NotNil(t, err)
	expectedError := "price invalid"
	assert.Equal(t, expectedError, err.Error())
	assert.Nil(t, plain)
}
