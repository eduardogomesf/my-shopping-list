package usecases

import "errors"

type UseCaseErrors struct {
	ErrUnfinishedShoppingList error
}

var useCaseErrors = UseCaseErrors{
	ErrUnfinishedShoppingList: errors.New("there is an unfinished shopping list with the given name"),
}

func GetUseCaseErrors() UseCaseErrors {
	return useCaseErrors
}
