package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePrice_PercentageDiscount(t *testing.T) {
	basePrice := 100.0
	discount := "p20"
	taxRate := 0.1

	finalPrice := calculatePrice(basePrice, discount, taxRate)
	expectedPrice := 100.0 * 0.8 * 1.1 // 20% скидка и 10% налог

	assert.Equal(t, expectedPrice, finalPrice)
}

func TestCalculatePrice_FixedDiscount(t *testing.T) {
	basePrice := 100.0
	discount := "10"
	taxRate := 0.1

	finalPrice := calculatePrice(basePrice, discount, taxRate)
	expectedPrice := 90.0 * 1.1 // Фиксированная скидка 10 и 10% налог

	assert.Equal(t, expectedPrice, finalPrice)
}

func TestCalculatePrice_MinimumLimit(t *testing.T) {
	basePrice := 100.0
	discount := "p60"
	taxRate := 0.1

	finalPrice := calculatePrice(basePrice, discount, taxRate)
	expectedPrice := 50.0 * 1.1 // 60% скидка, но со снижением до 50%

	assert.Equal(t, expectedPrice, finalPrice)
}

// Интеграционные тесты требуют поднятой инфраструктуры

func TestGetProductPriceEndpoint(t *testing.T) {
	// Для интеграционных тестов вы можете использовать http-тестирование
	// типо httptest.NewRequest или curl команды через Exec или другие инструменты.
}

func TestAddProductEndpoint(t *testing.T) {
	// Аналогично, протестируйте HTTP-запрос для добавления продукта.
}
