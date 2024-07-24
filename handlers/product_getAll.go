package handlers

import (
	"net/http"
	"project1/config"
	"project1/models"

	"github.com/labstack/echo/v4"
)

// GetAllProducts retrieves all products from the database
func GetAllProducts(c echo.Context) error {
    var products []models.Product
    query := `SELECT * FROM products`
    err := config.DB.Select(&products, query)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, products)
}
