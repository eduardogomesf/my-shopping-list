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

	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.PGHost, conf.PGPort, conf.PGUser, conf.PGPass, conf.PGDB)
	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.ShoppingListModel{})

	shoppingListRepository := repositories.NewShoppingListRepository(db)
	addShoppingListUseCase := usecases.NewAddShoppingListUseCase(shoppingListRepository, shoppingListRepository)
	addNewShoppingListHandler := handler.NewAddShoppingListHandler(*addShoppingListUseCase)
	ws.AddHandler("POST", "/shopping-lists", addNewShoppingListHandler.Handle)

	ws.Start()
}
