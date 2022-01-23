package routes

import (
	"project-e-commerces/constants"
	"project-e-commerces/delivery/controllers/carts"
	"project-e-commerces/delivery/controllers/categorys"
	"project-e-commerces/delivery/controllers/products"
	"project-e-commerces/delivery/controllers/transactions"
	"project-e-commerces/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController, crCtrl *carts.CartsController, tsCtrl *transactions.TransactionsController, cc *categorys.CategoryController, pc *products.ProductController) {

	// ---------------------------------------------------------------------
	// Login & Register
	// ---------------------------------------------------------------------
	e.POST("/users/register", uctrl.PostUserCtrl())
	e.POST("/users/login", uctrl.Login())

	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.GET("/users", uctrl.GetUsersCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.GET("/users/:id", uctrl.GetUserCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.PUT("/users/:id", uctrl.EditUserCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.DELETE("/users/:id", uctrl.DeleteUserCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	// ---------------------------------------------------------------------
	// CRUD Carts
	// ---------------------------------------------------------------------
	e.PUT("/carts/additem", crCtrl.PutItemIntoDetail_CartCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.GET("/carts", crCtrl.Gets(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.DELETE("/carts/delitem", crCtrl.DeleteItemFromDetail_CartCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.POST("/transactions/live", tsCtrl.PostProductsIntoTransactionCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.GET("/transactions/status", tsCtrl.GetStatus(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	// e.POST("/transactions/cart", tsCtrl.PostCartIntoTransactionCtrl(), middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	e.GET("/categorys", cc.GetAllCategory)
	e.GET("/categorys/:id", cc.GetCategoryByID)
	e.POST("/categorys", cc.CreateCategory)
	e.PUT("/categorys/:id", cc.UpdateCategory)
	e.DELETE("/categorys/:id", cc.DeleteCategory)

	e.GET("/products", pc.GetAllProduct)
	e.GET("/products/:id", pc.GetProductByID)
	e.GET("/products/stocks/:id", pc.GetHistoryStockProduct)
	e.POST("/products", pc.CreateProduct)
	e.POST("/products/stocks/:id", pc.UpdateStockProduct)
	e.PUT("/products/:id", pc.UpdateProduct)
	e.DELETE("/products/:id", pc.DeleteProduct)
}
