package usecases

import (
	"errors"
	"testing"

	"github.com/eduardogomesf/shopping/internal/dto"
	"github.com/eduardogomesf/shopping/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetActiveShoppingListByNameRepository struct {
	mock.Mock
}

func (m *MockGetActiveShoppingListByNameRepository) GetActiveByName(name string) (*entity.ShoppingList, error) {
	args := m.Called(name)
	return args.Get(0).(*entity.ShoppingList), args.Error(1)
}

type MockCreateShoppingListRepository struct {
	mock.Mock
}

func (m *MockCreateShoppingListRepository) Create(shoppingList *entity.ShoppingList) error {
	args := m.Called(shoppingList)
	return args.Error(0)
}

func TestAddShoppingListUseCase_Add_Success(t *testing.T) {
	mockGetActiveByNameRepo := new(MockGetActiveShoppingListByNameRepository)
	mockCreateRepo := new(MockCreateShoppingListRepository)
	useCase := AddShoppingListUseCase{
		GetActiveShoppingListByNameRepository: mockGetActiveByNameRepo,
		CreateShoppingListRepository:          mockCreateRepo,
	}

	mockGetActiveByNameRepo.On("GetActiveByName", "Groceries").Return((*entity.ShoppingList)(nil), nil)
	mockCreateRepo.On("Create", mock.AnythingOfType("*entity.ShoppingList")).Return(nil)

	err := useCase.Add(dto.AddShoppingListDTO{Name: "Groceries"})

	assert.NoError(t, err)
	mockGetActiveByNameRepo.AssertExpectations(t)
	mockCreateRepo.AssertExpectations(t)
}

func TestAddShoppingListUseCase_Add_ExistingShoppingListError(t *testing.T) {
	mockGetActiveByNameRepo := new(MockGetActiveShoppingListByNameRepository)
	mockCreateRepo := new(MockCreateShoppingListRepository)
	useCase := AddShoppingListUseCase{
		GetActiveShoppingListByNameRepository: mockGetActiveByNameRepo,
		CreateShoppingListRepository:          mockCreateRepo,
	}

	mockGetActiveByNameRepo.On("GetActiveByName", "Groceries").Return(&entity.ShoppingList{}, nil)

	err := useCase.Add(dto.AddShoppingListDTO{Name: "Groceries"})

	assert.EqualError(t, err, "there is an unfinished shopping list with the given name")
	mockGetActiveByNameRepo.AssertExpectations(t)
}

func TestAddShoppingListUseCase_Add_GetRepositoryError(t *testing.T) {
	mockGetActiveByNameRepo := new(MockGetActiveShoppingListByNameRepository)
	mockCreateRepo := new(MockCreateShoppingListRepository)
	useCase := AddShoppingListUseCase{
		GetActiveShoppingListByNameRepository: mockGetActiveByNameRepo,
		CreateShoppingListRepository:          mockCreateRepo,
	}

	mockGetActiveByNameRepo.On("GetActiveByName", "Groceries").Return((*entity.ShoppingList)(nil), errors.New("database error"))

	err := useCase.Add(dto.AddShoppingListDTO{Name: "Groceries"})

	assert.EqualError(t, err, "database error")
	mockGetActiveByNameRepo.AssertExpectations(t)
}

func TestAddShoppingListUseCase_Add_CreateRepositoryError(t *testing.T) {
	mockGetActiveByNameRepo := new(MockGetActiveShoppingListByNameRepository)
	mockCreateRepo := new(MockCreateShoppingListRepository)
	useCase := AddShoppingListUseCase{
		GetActiveShoppingListByNameRepository: mockGetActiveByNameRepo,
		CreateShoppingListRepository:          mockCreateRepo,
	}

	mockGetActiveByNameRepo.On("GetActiveByName", "Groceries").Return((*entity.ShoppingList)(nil), nil)
	mockCreateRepo.On("Create", mock.AnythingOfType("*entity.ShoppingList")).Return(errors.New("creation error"))

	err := useCase.Add(dto.AddShoppingListDTO{Name: "Groceries"})

	assert.EqualError(t, err, "creation error")
	mockGetActiveByNameRepo.AssertExpectations(t)
	mockCreateRepo.AssertExpectations(t)
}
