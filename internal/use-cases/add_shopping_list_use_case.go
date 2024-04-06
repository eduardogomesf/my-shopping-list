package usecase

import (
	"errors"
	"time"

	"github.com/eduardogomesf/shopping/internal/dto"
	"github.com/eduardogomesf/shopping/internal/entity"
	"github.com/google/uuid"
)

type AddShoppingListUseCase struct {
	GetActiveShoppingListByNameRepository GetActiveShoppingListByNameRepository
	CreateShoppingListRepository          CreateShoppingListRepository
}

func (asl *AddShoppingListUseCase) Add(data dto.AddShoppingListDTO) error {
	shoppingListByName, err := asl.GetActiveShoppingListByNameRepository.GetActiveByName(data.Name)

	if err != nil {
		return err
	}

	if shoppingListByName != nil {
		return errors.New("there is an unfinished shopping list with the given name")
	}

	shoppingList, err := entity.NewShoppingList(
		uuid.NewString(),
		data.Name,
		false,
		time.Now().UTC(),
		nil,
	)

	if err != nil {
		return err
	}

	err = asl.CreateShoppingListRepository.Create(shoppingList)

	if err != nil {
		return err
	}

	return nil
}
