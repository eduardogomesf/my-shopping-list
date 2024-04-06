package repositories

import (
	"errors"

	"github.com/eduardogomesf/shopping/internal/entity"
	"github.com/eduardogomesf/shopping/internal/infra/databases/models"
	"gorm.io/gorm"
)

type ShoppingListRepository struct {
	DbConnection *gorm.DB
}

func NewShoppingListRepository(dbConnection *gorm.DB) *ShoppingListRepository {
	return &ShoppingListRepository{
		DbConnection: dbConnection,
	}
}

func (slr *ShoppingListRepository) GetActiveByName(name string) (*entity.ShoppingList, error) {
	var shoppingList models.ShoppingListModel

	result := slr.DbConnection.Where("name = ?", name).Where("is_finished = ?", false).First(&shoppingList)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &entity.ShoppingList{
		ID:         shoppingList.ID,
		Name:       shoppingList.Name,
		CreatedAt:  shoppingList.CreatedAt,
		FinishedAt: shoppingList.FinishedAt,
		IsFinished: shoppingList.IsFinished,
	}, nil
}

func (slr *ShoppingListRepository) Create(sl *entity.ShoppingList) error {
	result := slr.DbConnection.Create(&models.ShoppingListModel{
		ID:         sl.ID,
		Name:       sl.Name,
		CreatedAt:  sl.CreatedAt,
		FinishedAt: sl.FinishedAt,
		IsFinished: sl.IsFinished,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
