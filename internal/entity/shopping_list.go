package entity

import (
	"errors"
	"time"
)

type ShoppingList struct {
	ID         string
	Name       string
	CreatedAt  time.Time
	FinishedAt *time.Time
	IsFinished bool
}

func NewShoppingList(id string, name string, isFinished bool, createdAt time.Time, finishedAt *time.Time) (*ShoppingList, error) {
	sl := &ShoppingList{
		ID:         id,
		Name:       name,
		IsFinished: isFinished,
		CreatedAt:  createdAt,
		FinishedAt: finishedAt,
	}

	err := sl.IsValid()

	if err != nil {
		return nil, err
	}

	return sl, nil
}

func (sl *ShoppingList) IsValid() error {
	if sl.ID == "" {
		return errors.New("invalid id")
	}

	if sl.Name == "" {
		return errors.New("invalid name")
	}

	return nil
}
