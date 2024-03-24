package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewShoppingListValid(t *testing.T) {
	id := "1"
	name := "Groceries"
	isFinished := false
	createdAt := time.Now()
	var finishedAt *time.Time

	sl, err := NewShoppingList(id, name, isFinished, createdAt, finishedAt)

	assert.Nil(t, err)
	assert.NotNil(t, sl)
	assert.Equal(t, id, sl.ID)
	assert.Equal(t, name, sl.Name)
	assert.Equal(t, isFinished, sl.IsFinished)
	assert.Equal(t, createdAt, sl.CreatedAt)
	assert.Equal(t, finishedAt, sl.FinishedAt)
}

func TestNewShoppingListInvalidID(t *testing.T) {
	id := "" // Invalid ID
	name := "Groceries"
	isFinished := false
	createdAt := time.Now()
	var finishedAt *time.Time

	sl, err := NewShoppingList(id, name, isFinished, createdAt, finishedAt)

	assert.Nil(t, sl)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid id", err.Error())
}

func TestNewShoppingListInvalidName(t *testing.T) {
	id := "1"
	name := "" // Invalid name
	isFinished := false
	createdAt := time.Now()
	var finishedAt *time.Time

	sl, err := NewShoppingList(id, name, isFinished, createdAt, finishedAt)

	assert.Nil(t, sl)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid name", err.Error())
}
