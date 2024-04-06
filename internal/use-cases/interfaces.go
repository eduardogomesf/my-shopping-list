package usecase

import "github.com/eduardogomesf/shopping/internal/entity"

type GetActiveShoppingListByNameRepository interface {
	GetActiveByName(name string) (*entity.ShoppingList, error)
}

type CreateShoppingListRepository interface {
	Create(sl *entity.ShoppingList) error
}
