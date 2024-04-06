package usecase

import (
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

	useCaseErrors := GetUseCaseErrors()

	if shoppingListByName != nil {
		return useCaseErrors.ErrUnfinishedShoppingList
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
