package repositories

import (
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

func (slr *ShoppingListRepository) Get(name string) (*entity.ShoppingList, error) {

}

func (slr *ShoppingListRepository) Create(sl *entity.ShoppingList) error {
	result := slr.DbConnection.Create(&models.ShoppingListModel{
		ID:         sl.ID,
		Name:       sl.Name,
		CreatedAt:  sl.CreatedAt,
		FinishedAt: sl.FinishedAt,
		IsFinished: sl.IsFinished,
	})

}
