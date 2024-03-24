package entity

import "errors"

type ListItem struct {
	ID             string
	Name           string
	Price          int16
	Quantity       int16
	Description    *string
	ShoppingListId string
}

func NewListItem(id string, name string, quantity int16, price int16, description *string, shoppingListId string) (*ListItem, error) {
	li := &ListItem{
		ID:             id,
		Name:           name,
		Quantity:       quantity,
		Price:          price,
		Description:    description,
		ShoppingListId: shoppingListId,
	}

	err := li.IsValid()

	if err != nil {
		return nil, err
	}

	return li, nil
}

func (li *ListItem) IsValid() error {
	if li.ID == "" {
		return errors.New("invalid id")
	}

	if li.Name == "" {
		return errors.New("invalid name")
	}

	if li.Quantity <= 0 {
		return errors.New("invalid quantity")
	}

	if li.Price <= 0 {
		return errors.New("invalid price")
	}

	if li.ShoppingListId == "" {
		return errors.New("invalid shopping list id")
	}

	return nil
}
