package handlers

import (
	"net/http"
	"project1/config"

	"github.com/labstack/echo/v4"
)

func DeleteProduct(c echo.Context) error {
    id := c.Param("id")
    query := `DELETE FROM products WHERE id=?`
    _, err := config.DB.Exec(query, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.NoContent(http.StatusNoContent)
}
