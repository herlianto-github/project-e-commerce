package main

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/delivery/controllers/carts"
	tempCategoryController "project-e-commerces/delivery/controllers/categorys"
	tempProductController "project-e-commerces/delivery/controllers/products"
	"project-e-commerces/delivery/controllers/transactions"
	"project-e-commerces/delivery/controllers/users"
	"project-e-commerces/delivery/routes"
	cartsRepo "project-e-commerces/repository/carts"
	tempCategoryRepo "project-e-commerces/repository/categorys"
	tempProductRepo "project-e-commerces/repository/products"
	transactionsRepo "project-e-commerces/repository/transactions"
	usersRepo "project-e-commerces/repository/users"
	"project-e-commerces/utils"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c coreapi.Client

func main() {

	midtrans.ServerKey = "SB-Mid-server-WBQoXNegZ5veTRfQsX3WOGFq"
	midtrans.ClientKey = "SB-Mid-client-lbfJ_9e_8nsyvWWS"
	midtrans.Environment = midtrans.Sandbox

	c.New(midtrans.ServerKey, midtrans.Environment)

	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()
	usersRepo := usersRepo.NewUsersRepo(db)
	usersCtrl := users.NewUsersControllers(usersRepo)

	cartsRepo := cartsRepo.NewCartsRepo(db)
	cartsCtrl := carts.NewCartsControllers(cartsRepo)

	transactionsRepo := transactionsRepo.NewTransactionsRepo(db)
	transactionsCtrl := transactions.NewTransactionsControllers(transactionsRepo)

	categoryRepo := tempCategoryRepo.NewCategoryRepo(db)
	categoryController := tempCategoryController.NewCategoryControllers(categoryRepo)

	productRepo := tempProductRepo.NewProductRepo(db)
	productController := tempProductController.NewProductControllers(productRepo)

	routes.RegisterPath(e, usersCtrl, cartsCtrl, transactionsCtrl, categoryController, productController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
