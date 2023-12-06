package http

import (
	"net/http"
	
	"github.com/labstack/echo/v4"

	"grpc-example/application/grpc"
	"grpc-example/model"
)

type webServer struct {

}

func NewWebServer() *webServer {
	return &webServer{}
}

func (w webServer) Serve(port string) {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.POST("/products", w.createProduct)
	e.Start(port)
}

func (w webServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, grpc.ProductList)
}

func (w webServer) createProduct(c echo.Context) error {
	product := model.NewProduct()
	if err := c.Bind(product); err != nil {
		return err
	}
	grpc.ProductList.Add(product)
	return c.JSON(http.StatusOK, product)
}