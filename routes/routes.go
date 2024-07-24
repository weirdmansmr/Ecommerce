package routes

import (
	"net/http"

	"project1/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.GET("/products", handlers.GetAllProducts)
    e.POST("/products", handlers.CreateProduct)
    e.PUT("/products/:id", handlers.UpdateProduct)
    e.DELETE("/products/:id", handlers.DeleteProduct)
}
