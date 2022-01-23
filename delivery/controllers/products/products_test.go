package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/delivery/pagination"
	"project-e-commerces/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProduct(t *testing.T) {
	t.Run("1-success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockProductRepository{})
		productController.GetAllProduct(context)

		response := pagination.ProductPagination{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// data := response.Data
		// fmt.Println(res)

		// name := data[0].Name

		product := response.Rows

		assert.Equal(t, "Product Alpha", product[0].Name)
		// assert.Equal(t, "success get all product", response.Message)
	})

	t.Run("2-error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.GetAllProduct(context)

		response := pagination.ProductPagination{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		product := response.Rows

		assert.Equal(t, []entities.Product([]entities.Product(nil)), product)

		// assert.Equal(t, []entities.Product(nil), response.Data)

		// assert.Equal(t, "no products data found", response.Message)
	})
}

func TestGetProductByID(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductControllers(mockProductRepository{})
		productController.GetProductByID(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data[0]

		name := data.Name

		assert.Equal(t, name, "Product Alpha")
		assert.Equal(t, "success get product", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/?", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)

		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("2")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.GetProductByID(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Product([]entities.Product(nil)), response.Data)
		assert.Equal(t, "product not found", response.Message)
	})
}

func TestGetStockHistoryProduct(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products/stock/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductControllers(mockProductRepository{})
		productController.GetHistoryStockProduct(context)

		response := GetStockProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		qty := response.Data[0].Qty

		assert.Equal(t, 1, qty)
		assert.Equal(t, "success get history product", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products/stock/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.GetHistoryStockProduct(context)

		response := GetStockProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Stock(nil), response.Data)
		assert.Equal(t, "history product not found", response.Message)
	})
}

func TestUpdateStockProduct(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Stock{
			Product_id: 1,
			Qty:        1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products/stock/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductControllers(mockProductRepository{})
		productController.UpdateStockProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data

		stock := data[0].Stock

		assert.Equal(t, 2, stock)
		assert.Equal(t, "success update stock product", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Stock{
			Product_id: 2,
			Qty:        1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products/stocks")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.UpdateStockProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Product([]entities.Product(nil)), response.Data)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Product{
			Name:        "Product Alpha",
			Stock:       1,
			Price:       10000,
			Category_id: 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockProductRepository{})
		productController.CreateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data

		name := data[0].Name

		assert.Equal(t, "Product Alpha", name)
		assert.Equal(t, "success create product", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name": "Product Alpha",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.CreateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "can't create product", response.Message)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Product Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductControllers(mockProductRepository{})
		productController.UpdateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data

		name := data[0].Name

		assert.Equal(t, "Product Alpha new", name)
		assert.Equal(t, "success update product", response.Message)
	})
	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Product Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.UpdateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Product([]entities.Product(nil)), response.Data)
		assert.Equal(t, "can't update product", response.Message)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductControllers(mockProductRepository{})
		productController.DeleteProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "success delete product", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("10")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.DeleteProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "can't delete product", response.Message)
	})
}

func TestExportPDF(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/export")

		productController := NewProductControllers(mockProductRepository{})
		productController.ExportPDF(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "success export pdf", "success export pdf")
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/export")

		productController := NewProductControllers(mockProductRepository{})
		productController.ExportPDF(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response.Message)

		assert.Equal(t, "can't export pdf", "can't export pdf")
	})
}

type mockProductRepository struct{}

func (m mockProductRepository) GetAllProduct() ([]entities.Product, error) {
	return []entities.Product{
		{ID: 1, Name: "Product Alpha"},
	}, nil
}

func (m mockProductRepository) GetProductByID(product_id int) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha"}, nil
}

func (m mockProductRepository) GetHistoryStockProduct(product_id int) ([]entities.Stock, error) {
	return []entities.Stock{
		{Product_id: 1, Qty: 1}}, nil
}

func (m mockProductRepository) CreateProduct(entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha"}, nil
}

func (m mockProductRepository) UpdateProduct(product_id int, product entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha new"}, nil
}

func (m mockProductRepository) UpdateStockProduct(product_id, qty int) (entities.Product, error) {
	return entities.Product{
		ID: 1, Stock: 2}, nil
}

func (m mockProductRepository) DeleteProduct(product_id int) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha"}, nil
}

func (m mockProductRepository) ProductPagination(Pagination pagination.ProductPagination) (interface{}, int, error) {
	return pagination.ProductPagination{
		Limit: 0, Page: 0, TotalRows: 5, FirstPage: "/products?limit=0&page=0", LastPage: "/products?limit=0&page=0",
		PreviousPage: "", NextPage: "/products?limit=0&page=1", FromRow: 1, ToRow: 5,
		Rows: []entities.Product{
			{Name: "Product Alpha", Price: 1000, Stock: 1, Category_id: 1},
			{Name: "Product Beta", Price: 2000, Stock: 2, Category_id: 2},
			{Name: "Product Cherry", Price: 3000, Stock: 3, Category_id: 3},
			{Name: "Product Delta", Price: 4000, Stock: 4, Category_id: 4},
			{Name: "Product Echo", Price: 5000, Stock: 5, Category_id: 5},
		},
	}, 5, nil
}

type mockFalseProductRepository struct{}

func (m mockFalseProductRepository) GetAllProduct() ([]entities.Product, error) {
	return nil, errors.New("no data")
}

func (m mockFalseProductRepository) GetProductByID(id int) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("can't get product")
}

func (m mockFalseProductRepository) GetHistoryStockProduct(product_id int) ([]entities.Stock, error) {
	return []entities.Stock{
		{Product_id: 0, Qty: 0}}, nil
}

func (m mockFalseProductRepository) CreateProduct(entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("error create product")
}

func (m mockFalseProductRepository) UpdateProduct(product_id int, product entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("error update product")
}

func (m mockFalseProductRepository) UpdateStockProduct(product_id, qty int) (entities.Product, error) {
	return entities.Product{
		ID: 0, Stock: 0}, errors.New("can't update stock product")
}

func (m mockFalseProductRepository) DeleteProduct(product_id int) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("error delete product")
}

func (m mockFalseProductRepository) ProductPagination(Pagination pagination.ProductPagination) (interface{}, int, error) {
	return pagination.ProductPagination{Limit: 0, Page: 0}, 0, errors.New("can't get products")
}
