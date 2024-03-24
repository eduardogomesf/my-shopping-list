package usecases

import "github.com/eduardogomesf/shopping/internal/entity"

type GetActiveShoppingListByNameRepository interface {
	Get(name string) (*entity.ShoppingList, error)
}

type CreateShoppingListRepository interface {
	Create(*entity.ShoppingList) error
}
