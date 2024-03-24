package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListItemValid(t *testing.T) {
	id := "123"
	name := "Bag of rice"
	quantity := int16(2)
	price := float32(5)
	description := ""
	shoppingListId := "321"

	li, err := NewListItem(
		id,
		name,
		quantity,
		price,
		&description,
		shoppingListId,
	)

	assert.Nil(t, err)
	assert.NotNil(t, li)
	assert.Equal(t, id, li.ID)
	assert.Equal(t, name, li.Name)
	assert.Equal(t, quantity, li.Quantity)
	assert.Equal(t, price, li.Price)
	assert.Equal(t, description, *li.Description)
	assert.Equal(t, shoppingListId, li.ShoppingListId)
}

func TestNewListItemInvalidID(t *testing.T) {
	id := ""
	name := "Bag of rice"
	quantity := int16(2)
	price := float32(5)
	description := ""
	shoppingListId := "321"

	li, err := NewListItem(
		id,
		name,
		quantity,
		price,
		&description,
		shoppingListId,
	)

	assert.Nil(t, li)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid id", err.Error())
}

func TestNewListItemInvalidName(t *testing.T) {
	id := "123"
	name := ""
	quantity := int16(2)
	price := float32(5)
	description := ""
	shoppingListId := "321"

	li, err := NewListItem(
		id,
		name,
		quantity,
		price,
		&description,
		shoppingListId,
	)

	assert.Nil(t, li)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid name", err.Error())
}

func TestNewListItemInvalidQuantity(t *testing.T) {
	id := "123"
	name := "Bag of rice"
	quantity := int16(0)
	price := float32(5)
	description := ""
	shoppingListId := "321"

	li, err := NewListItem(
		id,
		name,
		quantity,
		price,
		&description,
		shoppingListId,
	)

	assert.Nil(t, li)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid quantity", err.Error())
}

func TestNewListItemInvalidPrice(t *testing.T) {
	id := "123"
	name := "Bag of rice"
	quantity := int16(2)
	price := float32(0)
	description := ""
	shoppingListId := "321"

	li, err := NewListItem(
		id,
		name,
		quantity,
		price,
		&description,
		shoppingListId,
	)

	assert.Nil(t, li)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid price", err.Error())
}

func TestNewListItemInvalidShoppingListID(t *testing.T) {
	id := "123"
	name := "Bag of rice"
	quantity := int16(2)
	price := float32(5)
	description := ""
	shoppingListId := ""

	li, err := NewListItem(
		id,
		name,
		quantity,
		price,
		&description,
		shoppingListId,
	)

	assert.Nil(t, li)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid shopping list id", err.Error())
}
