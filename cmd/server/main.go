package main

import (
	"fmt"

	config "github.com/eduardogomesf/shopping/configs"
	"github.com/eduardogomesf/shopping/internal/infra/databases/models"
	"github.com/eduardogomesf/shopping/internal/infra/databases/repositories"
	webserver "github.com/eduardogomesf/shopping/internal/infra/web"
	handler "github.com/eduardogomesf/shopping/internal/infra/web/handlers"
	usecases "github.com/eduardogomesf/shopping/internal/use-cases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := config.LoadConfig(".")
	ws := webserver.NewWebServer(conf.APPPort)

	connectionUrl := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.ShoppingListModel{})

	shoppingListRepository := repositories.NewShoppingListRepository(db)
	addShoppingListUseCase := usecases.AddShoppingListUseCase{
		GetActiveShoppingListByNameRepository: shoppingListRepository,
		CreateShoppingListRepository:          shoppingListRepository,
	}
	addNewShoppingListHandler := handler.NewAddShoppingListHandler(addShoppingListUseCase)
	ws.AddHandler("POST", "/shopping-lists", addNewShoppingListHandler.Handle)

	ws.Start()
}
