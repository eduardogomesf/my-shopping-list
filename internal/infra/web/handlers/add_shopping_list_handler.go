package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eduardogomesf/shopping/internal/dto"
	usecase "github.com/eduardogomesf/shopping/internal/use-cases"
)

type AddShoppingListHandler struct {
	AddShoppingListUseCase usecase.AddShoppingListUseCase
}

func NewAddShoppingListHandler() *AddShoppingListHandler {
	return &AddShoppingListHandler{}
}

func (h *AddShoppingListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("error reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var data dto.AddShoppingListDTO

	json.Unmarshal(body, &data)

	err = h.AddShoppingListUseCase.Add(data)

	if err != nil {
		fmt.Println("error adding shopping list", err)

		usecaseErrors := usecase.GetUseCaseErrors()

		if err == usecaseErrors.ErrUnfinishedShoppingList {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
}
