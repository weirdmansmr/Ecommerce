package handlers

import (
	"net/http"
	"project1/config"
	"project1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateProduct(c echo.Context) error {
    id := c.Param("id")
		productID, err := strconv.Atoi(id)
		if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
    }
    product := new(models.Product)
    if err := c.Bind(product); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }
    product.ID = productID
    query := `UPDATE products SET name=:name, description=:description, price=:price, updated_at=NOW() WHERE id=:id`
    _, err = config.DB.NamedExec(query, product)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, product)
}
